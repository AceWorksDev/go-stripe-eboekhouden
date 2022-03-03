package service

import (
	"context"

	"github.com/aceworksdev/go-stripe-eboekhouden/internal/domain/customer"
	"github.com/aceworksdev/go-stripe-eboekhouden/internal/domain/hooks"
	"github.com/aceworksdev/go-stripe-eboekhouden/internal/domain/invoice"
)

type hooksService struct {
	customer customer.Servicer
	invoice  invoice.Servicer
}

func NewHooks() (hooks.Servicer, error) {
	return &hooksService{}, nil
}

func (service *hooksService) AddCustomer(item customer.Servicer) {
	service.customer = item
}

func (service *hooksService) AddInvoice(item invoice.Servicer) {
	service.invoice = item
}

func (service *hooksService) InvoiceCreate(ctx context.Context, item *invoice.Service) error {
	itemCustomer, err := service.customer.GetBasedOnStripeID(ctx, item.StripeCustomerID)
	if err != nil {
		return err
	}

	if err := service.invoice.Create(ctx, item, itemCustomer.Code); err != nil {
		return err
	}
	return nil
}

func (service *hooksService) InvoiceUpdate(ctx context.Context, item *invoice.Service) error {
	itemCustomer, err := service.customer.GetBasedOnStripeID(ctx, item.StripeCustomerID)
	if err != nil {
		return err
	}

	if err := service.invoice.Update(ctx, item, itemCustomer.Code); err != nil {
		return err
	}
	return nil
}

func (service *hooksService) CustomerCreate(ctx context.Context, item *customer.Service) error {
	if err := service.customer.Create(ctx, item); err != nil {
		return err
	}
	return nil
}
func (service *hooksService) CustomerUpdate(ctx context.Context, item *customer.Service) error {
	if err := service.customer.Update(ctx, item); err != nil {
		return err
	}
	return nil
}
