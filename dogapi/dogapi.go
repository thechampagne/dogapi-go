/*
 * Copyright 2022 XXIV
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package dogapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type breedsResponse struct {
	Status  string `json:"status"`
	Message map[string][]string `json:"message"`
}

func getRequest(endpoint string) (string, error) {
	response, err := http.Get(fmt.Sprintf("https://dog.ceo/api/%s", endpoint))
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// RandomImage
//
// DISPLAY SINGLE RANDOM IMAGE FROM ALL DOGS COLLECTION
//
// Returns a random dog image
func RandomImage() (string, error) {
	res, err := getRequest("breeds/image/random")
	if err != nil {
		return "", err
	}
	var data response
	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Something went wrong while reading json: %s", err))
	}
	if data.Status != "success" {
		return "", errors.New(data.Message)
	}
	return data.Message, nil
}

// MultipleRandomImages
//
// DISPLAY MULTIPLE RANDOM IMAGES FROM ALL DOGS COLLECTION
//
// "images_number" number of images
//
// "NOTE" ~ Max number returned is 50
//
// Return multiple random dog image
func MultipleRandomImages(imagesNumber int8) ([]string, error) {
	res, err := getRequest(fmt.Sprintf("breeds/image/random/%d", imagesNumber))
	if err != nil {
		return []string{}, err
	}
	var data map[string]interface{}
	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		return []string{}, errors.New(fmt.Sprintf("Something went wrong while reading json: %s", err))
	}
	if data["status"] != "success" {
		return []string{}, errors.New(fmt.Sprintf("%v",data["message"]))
	}
	var slice []string
	val := reflect.ValueOf(data["message"])
	for i := 0; i < val.Len(); i++ {
		slice = append(slice, fmt.Sprintf("%s", val.Index(i)))
	}
	return slice, nil
}

// RandomImageByBreed
//
// RANDOM IMAGE FROM A BREED COLLECTION
//
// "breed" breed name
//
// Returns a random dog image from a breed, e.g. hound
func RandomImageByBreed(breed string) (string, error) {
	res, err := getRequest(fmt.Sprintf("breed/%s/images/random", strings.TrimSpace(breed)))
	if err != nil {
		return "", err
	}
	var data response
	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Something went wrong while reading json: %s", err))
	}
	if data.Status != "success" {
		return "", errors.New(data.Message)
	}
	return data.Message, nil
}

// MultipleRandomImagesByBreed
//
// MULTIPLE IMAGES FROM A BREED COLLECTION
//
// "breed" breed name
// "images_number" number of images
//
// Return multiple random dog image from a breed, e.g. hound
func MultipleRandomImagesByBreed(breed string, imagesNumber int64) ([]string, error) {
	res, err := getRequest(fmt.Sprintf("breed/%s/images/random/%d", strings.TrimSpace(breed), imagesNumber))
	if err != nil {
		return []string{}, err
	}
	var data map[string]interface{}
	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		return []string{}, errors.New(fmt.Sprintf("Something went wrong while reading json: %s", err))
	}
	if data["status"] != "success" {
		return []string{}, errors.New(fmt.Sprintf("%v",data["message"]))
	}
	var slice []string
	val := reflect.ValueOf(data["message"])
	for i := 0; i < val.Len(); i++ {
		slice = append(slice, fmt.Sprintf("%s", val.Index(i)))
	}
	return slice, nil
}

// ImagesByBreed
//
// ALL IMAGES FROM A BREED COLLECTION
//
// "breed" breed name
//
// Returns an array of all the images from a breed, e.g. hound
func ImagesByBreed(breed string) ([]string, error) {
	res, err := getRequest(fmt.Sprintf("breed/%s/images", strings.TrimSpace(breed)))
	if err != nil {
		return []string{}, err
	}
	var data map[string]interface{}
	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		return []string{}, errors.New(fmt.Sprintf("Something went wrong while reading json: %s", err))
	}
	if data["status"] != "success" {
		return []string{}, errors.New(fmt.Sprintf("%v",data["message"]))
	}
	var slice []string
	val := reflect.ValueOf(data["message"])
	for i := 0; i < val.Len(); i++ {
		slice = append(slice, fmt.Sprintf("%s", val.Index(i)))
	}
	return slice, nil
}

// RandomImageBySubBreed
//
// SINGLE RANDOM IMAGE FROM A SUB BREED COLLECTION
//
// "breed" breed name
// "sub_breed" sub_breed name
//
// Returns a random dog image from a sub-breed, e.g. Afghan Hound
func RandomImageBySubBreed(breed string, subBreed string) (string, error) {
	res, err := getRequest(fmt.Sprintf("breed/%s/%s/images/random", strings.TrimSpace(breed), strings.TrimSpace(subBreed)))
	if err != nil {
		return "", err
	}
	var data response
	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Something went wrong while reading json: %s", err))
	}
	if data.Status != "success" {
		return "", errors.New(data.Message)
	}
	return data.Message, nil
}

// MultipleRandomImagesBySubBreed
//
// MULTIPLE IMAGES FROM A SUB-BREED COLLECTION
//
// "breed" breed name
// "sub_breed" sub_breed name
// "images_number" number of images
//
// Return multiple random dog images from a sub-breed, e.g. Afghan Hound
func MultipleRandomImagesBySubBreed(breed string, subBreed string, imagesNumber int64) ([]string, error) {
	res, err := getRequest(fmt.Sprintf("breed/%s/%s/images/random/%d", strings.TrimSpace(breed), strings.TrimSpace(subBreed), imagesNumber))
	if err != nil {
		return []string{}, err
	}
	var data map[string]interface{}
	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		return []string{}, errors.New(fmt.Sprintf("Something went wrong while reading json: %s", err))
	}
	if data["status"] != "success" {
		return []string{}, errors.New(fmt.Sprintf("%v",data["message"]))
	}
	var slice []string
	val := reflect.ValueOf(data["message"])
	for i := 0; i < val.Len(); i++ {
		slice = append(slice, fmt.Sprintf("%s", val.Index(i)))
	}
	return slice, nil
}

// ImagesBySubBreed
//
// LIST ALL SUB-BREED IMAGES
//
// "breed" breed name
// "sub_breed" sub_breed name
//
// Returns an array of all the images from the sub-breed
func ImagesBySubBreed(breed string, subBreed string) ([]string, error) {
	res, err := getRequest(fmt.Sprintf("breed/%s/%s/images", strings.TrimSpace(breed), strings.TrimSpace(subBreed)))
	if err != nil {
		return []string{}, err
	}
	var data map[string]interface{}
	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		return []string{}, errors.New(fmt.Sprintf("Something went wrong while reading json: %s", err))
	}
	if data["status"] != "success" {
		return []string{}, errors.New(fmt.Sprintf("%v",data["message"]))
	}
	var slice []string
	val := reflect.ValueOf(data["message"])
	for i := 0; i < val.Len(); i++ {
		slice = append(slice, fmt.Sprintf("%s", val.Index(i)))
	}
	return slice, nil
}

// BreedsList
//
// LIST ALL BREEDS
//
// Returns map of all the breeds as keys and sub-breeds as values if it has
func BreedsList() (map[string][]string, error) {
	res, err := getRequest("breeds/list/all")
	if err != nil {
		return map[string][]string{}, err
	}
	var data map[string]interface{}
	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		return map[string][]string{}, errors.New(fmt.Sprintf("Something went wrong while reading json: %s", err))
	}
	if data["status"] != "success" {
		return map[string][]string{}, errors.New(fmt.Sprintf("%v",data["message"]))
	}
	var br breedsResponse
	err = json.Unmarshal([]byte(res), &br)
	if err != nil {
		return map[string][]string{}, errors.New(fmt.Sprintf("Something went wrong while reading json: %s", err))
	}
	return br.Message, nil
}

// SubBreedsList
//
// LIST ALL SUB-BREEDS
//
// "breed" breed name
//
// Returns an array of all the sub-breeds from a breed if it has sub-breeds
func SubBreedsList(breed string) ([]string, error) {
	res, err := getRequest(fmt.Sprintf("breed/%s/list", strings.TrimSpace(breed)))
	if err != nil {
		return []string{}, err
	}
	var data map[string]interface{}
	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		return []string{}, errors.New(fmt.Sprintf("Something went wrong while reading json: %s", err))
	}
	if data["status"] != "success" {
		return []string{}, errors.New(fmt.Sprintf("%v",data["message"]))
	}
	var slice []string
	val := reflect.ValueOf(data["message"])
	for i := 0; i < val.Len(); i++ {
		slice = append(slice, fmt.Sprintf("%s", val.Index(i)))
	}
	if len(slice) == 0 {
		return []string{}, errors.New("the breed does not have sub-breeds")
	}
	return slice, nil
}