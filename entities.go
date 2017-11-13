package apiai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type EntityDescription struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Count   int    `json:"count"`
	Preview string `json:"preview"`
}

type Entry struct {
	Value    string   `json:"value"`
	Synonyms []string `json:"synonyms"`
}

type Entity struct {
	Id                 string  `json:"id"`
	Name               string  `json:"name"`
	Entries            []Entry `json:"entries"`
	IsEnum             bool    `json:"isEnum"`
	AutomatedExpansion bool    `json:"automatedExpansion"`
}

func (c *ApiClient) GetEntities() ([]EntityDescription, error) {
	req, err := http.NewRequest("GET", c.buildUrl("entities", nil), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.config.Token)

	// httpClient := http.DefaultClient
	httpClient := &http.Client{}

	if len(c.config.ProxyURL) > 0 {
		url, err := url.Parse(c.config.ProxyURL)
		if err == nil {
			httpClient.Transport = &http.Transport{Proxy: http.ProxyURL(url)}
		}
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var entities []EntityDescription
	switch resp.StatusCode {
	case http.StatusOK:
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&entities)
		if err != nil {
			return nil, err
		}
		return entities, nil
	default:
		return nil, fmt.Errorf("apiai: wops something happens because status code is %v", resp.StatusCode)
	}
}

func (c *ApiClient) GetEntity(idOrName string) (*Entity, error) {
	req, err := http.NewRequest("GET", c.buildUrl("entities/"+idOrName, nil), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.config.Token)

	// httpClient := http.DefaultClient
	httpClient := &http.Client{}

	if len(c.config.ProxyURL) > 0 {
		url, err := url.Parse(c.config.ProxyURL)
		if err == nil {
			httpClient.Transport = &http.Transport{Proxy: http.ProxyURL(url)}
		}
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var entity *Entity
	switch resp.StatusCode {
	case http.StatusOK:
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&entity)
		if err != nil {
			return nil, err
		}
		return entity, nil
	default:
		return nil, fmt.Errorf("apiai: wops something happens because status code is %v", resp.StatusCode)
	}
}

func (c *ApiClient) CreateEntity(entity Entity) (*CreationResponse, error) {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(entity)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.buildUrl("entities", nil), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+c.config.Token)

	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cr *CreationResponse
	switch resp.StatusCode {
	case http.StatusOK:
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&cr)
		if err != nil {
			return nil, err
		}
		return cr, nil
	default:
		return nil, fmt.Errorf("apiai: wops something happens because status code is %v", resp.StatusCode)
	}
}

func (c *ApiClient) AddEntries(idOrName string, entries []Entry) error {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(entries)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", c.buildUrl("entities/"+idOrName+"/entries", nil), body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+c.config.Token)

	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	default:
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("body: %s", body)
		return fmt.Errorf("apiai: wops something happens because status code is %v", resp.StatusCode)
	}
}

func (c *ApiClient) UpdateEntities(entities []Entity) error {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(entities)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", c.buildUrl("entities", nil), body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+c.config.Token)

	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf("apiai: wops something happens because status code is %v", resp.StatusCode)
	}
}

func (c *ApiClient) UpdateEntity(idOrName string, entity Entity) error {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(entity)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", c.buildUrl("entities/"+idOrName, nil), body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+c.config.Token)

	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf("apiai: wops something happens because status code is %v", resp.StatusCode)
	}
}

func (c *ApiClient) UpdateEntries(idOrName string, entries []Entry) error {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(entries)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", c.buildUrl("entities/"+idOrName+"/entries", nil), body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+c.config.Token)

	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	default:
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("body: %s", body)
		return fmt.Errorf("apiai: wops something happens because status code is %v", resp.StatusCode)
	}
}

func (c *ApiClient) DeleteEntity(idOrName string) error {
	req, err := http.NewRequest("DELETE", c.buildUrl("entities/"+idOrName, nil), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+c.config.Token)

	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf("apiai: wops something happens because status code is %v", resp.StatusCode)
	}
}

func (c *ApiClient) DeleteEntries(idOrName string, entries []string) error {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(entries)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("DELETE", c.buildUrl("entities/"+idOrName+"/entries", nil), body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+c.config.Token)

	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf("apiai: wops something happens because status code is %v", resp.StatusCode)
	}
}
