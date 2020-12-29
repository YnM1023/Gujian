// Package calculation implements the calculate rpc service.
package calculation

import (
	"context"
	pb "marinette/gujian/proto/calculation"
)

// Service implements Calculate rpc logics.
type Service struct {
	pb.UnimplementedCalculateServer
}

// AttributeCompare rpc service provides the calculation on dps gap between original attributes and new attributes.
// TODO(marinette@): implement the function.
func (s *Service) AttributeCompare(ctx context.Context, req *pb.AttributeCompareRequest) (*pb.AttributeCompareResponse, error) {
	return nil, nil
}
