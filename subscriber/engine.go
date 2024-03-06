package subscriber

import (
	"context"
	"github.com/yuisofull/food-delivery-app-with-go/common"
	"github.com/yuisofull/food-delivery-app-with-go/component/appctx"
	"github.com/yuisofull/food-delivery-app-with-go/component/asyncjob"
	"github.com/yuisofull/food-delivery-app-with-go/pubsub"
	"log"
)

type consumerJob struct {
	Title string
	Hld   func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	appCtx appctx.AppContext
}

func NewEngine(appCtx appctx.AppContext) *consumerEngine {
	return &consumerEngine{appCtx: appCtx}
}

func (engine *consumerEngine) Start() error {

	_ = engine.startSubTopic(
		common.TopicUserLikeRestaurant,
		true,
		IncreaseLikeCountAfterUserLikeRestaurant(engine.appCtx),
		PushNotificationAfterUserLikeRestaurant(engine.appCtx),
		EmitRealtimeAfterUserLikeRestaurant(engine.appCtx),
	)

	_ = engine.startSubTopic(
		common.TopicUserDislikeRestaurant,
		true,
		DecreaseLikeCountAfterUserDislikeRestaurant(engine.appCtx),
	)

	return nil
}

//type GroupJob interface {
//	Run(ctx context.Context) error
//}

func (engine *consumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, consumerJobs ...consumerJob) error {
	c, _ := engine.appCtx.GetPubSub().Subscribe(context.Background(), topic)

	for _, item := range consumerJobs {
		log.Println("setup consumer for: ", item.Title)
	}

	getJobHandler := func(job *consumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Println("running job for ", job.Title, ". Value: ", message.Data())
			return job.Hld(ctx, message)
		}
	}

	// linking the pub/sub system & async job system
	go func() {
		for {
			msg := <-c

			jobHldArr := make([]asyncjob.Job, len(consumerJobs))

			for i := range consumerJobs {
				jobHdl := getJobHandler(&consumerJobs[i], msg)
				jobHldArr[i] = asyncjob.NewJob(jobHdl)
			}

			group := asyncjob.NewGroup(isConcurrent, jobHldArr...)
			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}
