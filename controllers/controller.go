// Package controller Student API.
//
// the purpose of this application is to provide an application
// that is using go code to define an  Rest API
//
//     Schemes: http, https
//     Host: localhost:3000
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package controller

import (
	"Univercity/student"
	"Univercity/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var Client *mongo.Client

//home page
func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "university ")

}

// swagger:operation POST /students Student Poststudent
//
// Add new Student
//
// Returns new Student
//
// ---
// produces:
// - application/json
// parameters:
// - name: student
//   in: body
//   description: add Student data
//   required: true
//   schema:
//     "$ref": "#/definitions/Student"
// responses:
//   '200':
//     description:  New Student created
//     schema:
//       "$ref": "#/definitions/StudentStudent"
//   '409':
//     description: Conflict
//   '405':
//     description: Method Not Allowed
//   '403':
//     description: Forbidden

//Poststudent in DB
func Poststudent(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "New User creating ")

	fmt.Println("Show result ", r.Body)
	var student student.Student
	collection := utils.NewFunction()
	fmt.Println("Test", r.Body)
	ctx := r.Context()
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		log.Print(err)

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"message":"` + err.Error() + `"}"`))
		return
	}

	result, err := collection.InsertOne(ctx, student)
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"message":"` + err.Error() + `"}"`))
		return
	}
	fmt.Println(result)
	w.Header().Add("content-type", "appllication/json")

	json.NewEncoder(w).Encode(result)

}

// swagger:operation GET /students Student GetStudent
//
// Get Student
//
// Returns existing student
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: Student data
//     schema:
//      "$ref": "#/definitions/Student"

//GetStudent fetch all  from db
func GetStudent(w http.ResponseWriter, r *http.Request) {
	//var allstudent student.Students
	//var students []student.Student
	var allstudent []student.Student
	collection := utils.NewFunction()
	ctx := r.Context()
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return
	}

	for cur.Next(ctx) {
		var student student.Student
		err := cur.Decode(&student)
		if err != nil {
			return
		}
		//students = append(students, student)
		allstudent = append(allstudent, student)
	}
	//	allstudent.Students = students
	w.Header().Add("content-type", "appllication/json")

	json.NewEncoder(w).Encode(allstudent)

}

func GetStudentbyName(w http.ResponseWriter, r *http.Request) {
	//var students []Student

	// name := mux.Vars(r)["name"]

	ctx := r.Context()
	collection := utils.NewFunction()

	name := r.URL.Query().Get("name")
	city := r.URL.Query().Get("city")
	//year := r.URL.Query().Get("yearofaddmision")
	params := []primitive.M{}

	filter := primitive.M{}

	if name != "" {
		params = append(params, primitive.M{"name": name})
		filter = primitive.M{"name": name}
	}
	if city != "" {
		params = append(params, primitive.M{"city": city})
		filter = primitive.M{"city": city}
	}
	//if year != "" {
	//	params = append(params, primitive.M{"yearofaddmision": year})
	//	filter = primitive.M{"yearofaddmision": year}
	//}
	if len(params) > 1 {
		filter = primitive.M{"$and": params}
	}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return
	}
	students := []student.Student{}
	for cur.Next(ctx) {
		var std student.Student
		err := cur.Decode(&std)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		//fmt.Printf("\n%+v", std)
		students = append(students, std)
		// json.NewEncoder(w).Encode(student)
	}
	json.NewEncoder(w).Encode(students)

}

// swagger:operation GET /students/{id} Student GetStudentbyid
//
// Get Student
//
// Returns existing Student filtered by id
//
// ---
// produces:
// - application/json
// parameters:
//  - name: id
//    type: string
//    in: path
//    required: true
// responses:
//   '200':
//     description: Student data
//     schema:
//      "$ref": "#/definitions/Student"
//   '405':
//     description: Method Not Allowed
//   '403':
//     description: Forbidden

//GetStudentbyid fetch Student from db
func GetStudentbyid(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "GET BY ID STUDENT ")

	w.Header().Add("content-type", "application/json")
	var student student.Student
	id := mux.Vars(r)["id"]

	ID, err := primitive.ObjectIDFromHex(id)
	log.Println(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"` + err.Error() + `"}"`))
		log.Println(err)
		return

	}

	filter := primitive.M{"_id": ID}

	collection := utils.NewFunction()
	ctx := r.Context()
	err = collection.FindOne(ctx, filter).Decode(&student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}"`))
		return
	}
	w.Header().Add("content-type", "appllication/json")
	json.NewEncoder(w).Encode(student)

}

// swagger:operation DELETE /students/{id} Student StudentDeletestudent
//
// Delete  student
//
// Delete existing Student filtered by id
//
// ---
//
// parameters:
//  - name: id
//    type: string
//    in: path
//    required: true
// responses:
//   '410':
//     description: delete Student sucessfully

//Deletestudent fetch Student from db
func Deletestudent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User ")
	id := mux.Vars(r)["id"]

	ctx := r.Context()

	collection := utils.NewFunction()

	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := primitive.M{"_id": ID}
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		w.WriteHeader(http.StatusGone)
		//w.Write([]byte(`{"message":"` + err.Error() + `"}"`))
		return
		//log.Fatal(err)
	}
	// log.Println(result)
	// w.Header().Add("content-type", "appllication/json")

	//json.NewEncoder(w).Encode(http.StatusGone)

	w.WriteHeader(http.StatusGone)
}

// swagger:operation PUT /students/{id} Student Updatestudent
//
// Update Student
//
// Update existing Student filtered by id
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - name: id
//   type: string
//   in: path
//   required: true
// - name: name
//   in: body
//   description: add Student data
//   required: true
//   schema:
//     "$ref": "#/definitions/Student"
// responses:
//   '200':
//     description: Student response
//     schema:
//       "$ref": "#/definitions/Student"

//updatestudent----
func Updatestudent(w http.ResponseWriter, r *http.Request) {

	collection := utils.NewFunction()
	ctx := r.Context()

	var params = mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		return
	}
	var student student.Student
	err = json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		return
	}
	filter := bson.M{"_id": id}

	update := primitive.M{
		"name":            student.Name,
		"city":            student.City,
		"country":         student.Country,
		"course":          student.Course,
		"yearofadmission": student.YearOfAdmission}

	err = collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&student)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"` + err.Error() + `"}"`))
		return

	}
	log.Println(student)
	w.Header().Add("content-type", "appllication/json")

	json.NewEncoder(w).Encode(student)

}

// swagger:operation PATCH /students/{id} Student Updatestudent_patch
//
//  Parratially Update Student
//
// Patch existing Student filtered by id
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - name: id
//   type: string
//   in: path
//   required: true
// - name: name
//   in: body
//   description: add Student data
//   required: true
//   schema:
//     "$ref": "#/definitions/Student"
// responses:
//   '200':
//     description: Student response
//     schema:
//       "$ref": "#/definitions/Student"

//Updatestudent_patch----
func Updatestudent_patch(w http.ResponseWriter, r *http.Request) {

	collection := utils.NewFunction()
	ctx := r.Context()

	var params = mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return

	}

	var student student.Student
	filter := primitive.M{"_id": id}

	var update map[string]string

	err = json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		return
	}
	err = collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}).Decode(&student)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"` + err.Error() + `"}"`))
		return

	}
	log.Println(student)
	w.Header().Add("content-type", "appllication/json")

	json.NewEncoder(w).Encode(student)

}
