/**
 * SPDX-License-Identifier: Apache-2.0
 * © Copyright 2023 Hewlett Packard Enterprise Development LP
 */
package nasMsgHandler

import (
	"encoding/hex"
	"errors"
	"my5G-RANTester/test/amf/context"
	"strings"

	"github.com/free5gc/nas"

	log "github.com/sirupsen/logrus"
)

func AuthenticationResponse(nasMsg *nas.Message, ue *context.UEContext) error {
	if nasMsg.AuthenticationResponse.AuthenticationResponseParameter == nil {
		return errors.New("AuthenticationResponseParameter is nil")
	}
	resStarb := nasMsg.AuthenticationResponse.AuthenticationResponseParameter.GetRES()
	resStar := hex.EncodeToString(resStarb[:])

	xresStar := ue.GetSecurityContext().GetXresStar()

	if strings.EqualFold(resStar, xresStar) {
		log.Info("[AMF] 5G AKA confirmation succeeded")
		ue.DerivateKamf()
		return nil
	} else {
		return errors.New(("5G AKA confirmation failed, expected res* " + xresStar + " but got " + resStar))
	}
}
