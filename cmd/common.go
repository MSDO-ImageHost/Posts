package main



// Check for required header fields
tokenString, tokenPresent := msg.Headers["jwt"].(string)
if !tokenPresent {
	headers["status_code"] = http.StatusUnauthorized
	headers["status_code_msg"] = http.StatusText(http.StatusUnauthorized)
	log.Println(_LOG_TAG, "Rejected request with correlation id", msg.CorrelationId)
	if err := PublicateResponse(handleConf, msg, headers, nil, false, start); err != nil {
		log.Fatal(_LOG_TAG, "Failed process response to", msg.CorrelationId, err)
	}
	continue
}

// Parse JWT
token, err := auth.Parse(tokenString)
if err != nil || token == nil {
	headers["status_code"] = http.StatusUnauthorized
	headers["status_code_msg"] = http.StatusText(http.StatusUnauthorized)
	log.Println(_LOG_TAG, "Rejected request with correlation id", msg.CorrelationId)
	if err := PublicateResponse(handleConf, msg, headers, nil, false, start); err != nil {
		log.Fatal(_LOG_TAG, "Failed process response to", msg.CorrelationId, err)
	}
	continue
}