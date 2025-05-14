package transaction

import (
	"net/url"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func CreateQuery(q url.Values) (bson.D, error){
	transactionId := q.Get("transactionId")
	query := bson.D{}
	if transactionId != "" {	
		objectId, err := primitive.ObjectIDFromHex(transactionId)
		if err != nil {
			return nil, err
		}
		query = append(query, bson.E{Key: "_id", Value: objectId})
		return query, nil
	} 
	

	userId := q.Get("userId")
	if userId != "" {
			userObjectId, err := primitive.ObjectIDFromHex(userId)
			if err != nil {
				return nil, err
			}
			query = append(query, bson.E{Key: "userId", Value: userObjectId})
	}



	category := q.Get("category")
	query = AddStringToQuery("category", category, query)

	name := q.Get("name")
	query = AddStringToQuery("name", name, query)


	// convert to dateTimes
	dateBefore := q.Get("dateBefore")
	query = AddDateTimeToQuery("date", dateBefore, "$lte", query)
	dateAfter := q.Get("dateAfter")
	query = AddDateTimeToQuery("date", dateAfter, "$gte", query)

	// convert to int
	amountGreater := q.Get("amountGreater")
	query = AddIntToQuery("amount", amountGreater, "$gte", query)
	amountLesser 	:= q.Get("amountLesser")
	query = AddIntToQuery("amount", amountLesser, "$lte", query)
	return query, nil
}


func AddStringToQuery(field string, value string, query bson.D) (bson.D) {
	if value == "" {
		return query
	}
	query = append(query, bson.E{Key: field, Value: value})
	return query
}


func AddDateTimeToQuery(field string, value string, comp string, query bson.D) (bson.D) {
	if value == "" {
		return query
	}
	// yyyy-mm-dd
	date, err := time.Parse("2006-01-02", value)
	if err != nil {
		return query
	}

	primDate := primitive.NewDateTimeFromTime(date)
	query = append(query, bson.E{
		Key: field, 
		Value: bson.D{
			{Key: comp, Value: primDate},
		},
	})
	return query
}

func AddIntToQuery(field string, value string, comp string, query bson.D) (bson.D) {
	if value == "" {
		return query
	}

	num, err := strconv.Atoi(value)
	if err != nil || num < 0 {
		return query
	}
	query = append(query, bson.E{
		Key: field, 
		Value: bson.D{
			{Key: comp, Value: num},
		},
	})
	return query
}