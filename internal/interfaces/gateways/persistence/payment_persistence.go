package persistence

import (
	"context"
	"log/slog"
	"payments/internal/domain/entity"
	"time"

	"payments/pkg/random"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentRepository struct {
	db         *mongo.Client
	collection *mongo.Collection
}

func NewPaymentRepository(db *mongo.Client) PaymentRepository {
	collection := db.Database("payments").Collection("payments")

	return PaymentRepository{db: db, collection: collection}
}

func (gr PaymentRepository) Get(identifier string) (*entity.Payment, bool, error) {
	payment := entity.Payment{}

	filter := bson.M{
		"identifier": identifier,
	}

	err := gr.collection.FindOne(context.Background(), filter).Decode(&payment)

	if err != nil && err.Error() == "mongo: no documents in result" {
		return &payment, false, nil
	} else if err != nil {
		slog.Error("Error on retrieve mongodb result", "error", err)
	}

	slog.Debug("PaymentRepository#Get",
		"error", err,
		"result", payment)

	return &payment, true, err
}

func (pr PaymentRepository) Create(p *entity.Payment) error {
	p.ID = random.Hex(10)
	p.CreatedAt = time.Now().UTC().Format(time.RFC3339)
	p.State = "waiting"

	// Insert the item into the collection
	insertResult, err := pr.collection.InsertOne(context.Background(), p)
	if err != nil {
		slog.Error("Error on insert mongodb result", "error", err)
	}

	slog.Debug("PaymentRepositoryPersistence#Create",
		"insertResult", insertResult,
		"error", err,
	)

	return err
}

func (pr PaymentRepository) SetApproved(paymentID, externalServiceIdentifier, newState string) error {
	filterFields := bson.M{"id": paymentID}

	updateFields := bson.M{"$set": bson.M{
		"status":                      newState,
		"external_service_identifier": externalServiceIdentifier}}

	_, err := pr.collection.UpdateOne(context.Background(), filterFields, updateFields)

	if err != nil {
		slog.Error("Error updating status",
			"paymentID", paymentID,
			"newState", newState,
			"error", err)
	}

	return err
}

func (pr PaymentRepository) SetError(paymentID string, err error) error {
	filter := bson.M{"id": paymentID}

	update := bson.M{"$set": bson.M{"error": err.Error(), "state": "error"}}

	_, updateerr := pr.collection.UpdateOne(context.Background(), filter, update)

	if updateerr != nil {
		slog.Error("Error updating status",
			"paymentID", paymentID,
			"err", err,
			"updateerr", updateerr)
	}

	return err
}
