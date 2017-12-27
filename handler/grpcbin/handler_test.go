package grpcbinhandler

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	pb "github.com/moul/pb/grpcbin/go-grpc"
)

func TestHandlers(t *testing.T) {
	Convey("Testing Handlers{}", t, func() {
		s := Handlers{}
		ctx := context.Background()

		Convey("Testing Index()", func() {
			resp, err := s.Index(ctx, &pb.EmptyMessage{})
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeNil)
			So(len(resp.Endpoints) > 0, ShouldBeTrue)
			So(len(resp.Endpoints[0].Path) > 0, ShouldBeTrue)
		})

		Convey("Testing Empty()", func() {
			resp, err := s.Empty(ctx, &pb.EmptyMessage{})
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeNil)
		})

		Convey("Testing DummyUnary()", func() {
			req := &pb.DummyMessage{
				FString: "hello world",
				FInt32:  42,
			}
			resp, err := s.DummyUnary(ctx, req)
			So(err, ShouldBeNil)
			So(resp.FString, ShouldEqual, "hello world")
			So(resp.FInt32, ShouldEqual, 42)
			So(resp, ShouldResemble, req)
		})

		Convey("Testing RandomError()", func() {
			_, err := s.RandomError(ctx, &pb.EmptyMessage{})
			So(err, ShouldNotBeNil)
		})

		Convey("Testing SpecificError()", func() {
			_, err := s.SpecificError(ctx, &pb.SpecificErrorRequest{
				Code: 1,
			})
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "rpc error: code = Canceled desc = Canceled")

			_, err = s.SpecificError(ctx, &pb.SpecificErrorRequest{
				Code: 2,
			})
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "rpc error: code = Unknown desc = Unknown")

			_, err = s.SpecificError(ctx, &pb.SpecificErrorRequest{
				Code:   3,
				Reason: "lorem ipsum",
			})
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "rpc error: code = InvalidArgument desc = lorem ipsum")
		})

		// Convey("Testing DummyBidirectionalStreamStream()", func() {})
		// Convey("Testing DummyClientStream()", func() {})
		// Convey("Testing DummyServerStream()", func() {})
		// Convey("Testing HeadersUnary()", func() {})
		// Convey("Testing NoResponseUnary()", func() {})
	})
}
