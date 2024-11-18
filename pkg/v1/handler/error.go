package handler

import (
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorResponse(field string, err error) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: err.Error(),
	}
}

func invalidArgumentError(violation []*errdetails.BadRequest_FieldViolation) error {
	badRequest := &errdetails.BadRequest{FieldViolations: violation}
	statusInvalid := status.New(codes.InvalidArgument, "invalid parameters")
	statuDetails, err := statusInvalid.WithDetails(badRequest)
	if err != nil {
		return statusInvalid.Err()
	}
	return statuDetails.Err()
}
func ErrorResponses(err error) []*errdetails.BadRequest_FieldViolation {
	var details []*errdetails.BadRequest_FieldViolation
	if ve, ok := err.(*protovalidate.ValidationError); ok {
		for _, violation := range ve.Violations {
			details = append(details, &errdetails.BadRequest_FieldViolation{
				Field:       *violation.FieldPath,
				Description: *violation.Message,
			})
		}

	}
	return details
}
