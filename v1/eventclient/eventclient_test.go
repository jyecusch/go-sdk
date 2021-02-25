package eventclient

import (
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	v1 "github.com/nitrictech/go-sdk/interfaces/nitric/v1"
	mock_v1 "github.com/nitrictech/go-sdk/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/protobuf/types/known/structpb"
)

var _ = Describe("Eventclient", func() {
	ctrl := gomock.NewController(GinkgoT())

	When("GetTopics", func() {
		When("Topics are available", func() {
			It("Should return the topics", func() {
				mockTopicClient := mock_v1.NewMockTopicClient(ctrl)

				By("Calling GetTopics")
				mockTopicClient.EXPECT().List(gomock.Any(), &v1.TopicListRequest{}).Return(&v1.TopicListResponse{
					Topics: []*v1.NitricTopic{{Name: "test-topic"}},
				}, nil)

				client := NewWithClient(nil, mockTopicClient)
				topics, err := client.GetTopics()

				By("Not returning an error")
				Expect(err).ShouldNot(HaveOccurred())

				By("Returning the topics")
				Expect(topics).To(Equal([]Topic{
					&NitricTopic{
						name: "test-topic",
					},
				}))
			})
		})

		When("An error is returned from the gRPC client", func() {
			It("Should return an error", func() {
				mockTopicClient := mock_v1.NewMockTopicClient(ctrl)

				By("Calling GetTopics")
				mockTopicClient.EXPECT().List(gomock.Any(), &v1.TopicListRequest{}).Return(nil, fmt.Errorf("mock error"))

				client := NewWithClient(nil, mockTopicClient)
				_, err := client.GetTopics()

				By("Not returning an error")
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	When("Publish", func() {
		When("The topic exists", func() {
			It("Should publish the event", func() {
				mockEventClient := mock_v1.NewMockEventClient(ctrl)

				By("Calling GetTopics")

				payload := map[string]interface{}{
					"test": "content",
				}

				payloadStruct, _ := structpb.NewStruct(payload)

				mockEventClient.EXPECT().Publish(gomock.Any(), &v1.EventPublishRequest{
					Topic: "test-topic",
					Event: &v1.NitricEvent{
						RequestId:   "abc123",
						PayloadType: "test-payload-type",
						Payload:     payloadStruct,
					},
				}).Return(&v1.EventPublishResponse{}, nil)

				client := NewWithClient(mockEventClient, nil)
				topicName := "test-topic"
				payloadType := "test-payload-type"
				requestId := "abc123"
				// FIXME: This interface doesn't match the others.
				reqId, err := client.Publish(PublishOptions{
					TopicName: &topicName,
					Event: &Event{
						Payload:     &payload,
						PayloadType: &payloadType,
						RequestId:   &requestId,
					},
				})

				By("Not returning an error")
				Expect(err).ShouldNot(HaveOccurred())

				By("Returning the request id")
				Expect(*reqId).To(Equal("abc123"))
			})

			When("No request id is specified", func() {
				It("Should publish the event", func() {
					mockEventClient := mock_v1.NewMockEventClient(ctrl)

					By("Calling GetTopics")
					payload := map[string]interface{}{
						"test": "content",
					}

					mockEventClient.EXPECT().Publish(gomock.Any(), gomock.Any()).Return(&v1.EventPublishResponse{}, nil)

					client := NewWithClient(mockEventClient, nil)
					topicName := "test-topic"
					payloadType := "test-payload-type"
					reqId, err := client.Publish(PublishOptions{
						TopicName: &topicName,
						Event: &Event{
							Payload:     &payload,
							PayloadType: &payloadType,
						},
					})

					By("Not returning an error")
					Expect(err).ShouldNot(HaveOccurred())

					By("Returning the request id")
					uuid.MustParse(*reqId)
				})
			})
		})

		When("An error is returned from the gRPC client", func() {
			It("Should return an error", func() {
				mockEventClient := mock_v1.NewMockEventClient(ctrl)

				By("Calling GetTopics")
				payload := map[string]interface{}{
					"test": "content",
				}

				mockEventClient.EXPECT().Publish(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("mock error"))

				client := NewWithClient(mockEventClient, nil)
				topicName := "test-topic"
				payloadType := "test-payload-type"
				requestID := "abc123"
				_, err := client.Publish(PublishOptions{
					TopicName: &topicName,
					Event: &Event{
						Payload:     &payload,
						PayloadType: &payloadType,
						RequestId:   &requestID,
					},
				})

				By("Returning an error")
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	When("Topic", func() {
		When("A topic has been instantiated", func() {
			topic := NitricTopic{
				name: "test-topic",
			}

			It("should return its name", func() {
				Expect(topic.GetName()).To(Equal("test-topic"))
			})

			It("should be printable", func() {
				Expect(topic.String()).To(Equal("test-topic"))
			})
		})
	})
})
