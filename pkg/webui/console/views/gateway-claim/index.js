// Copyright © 2021 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'
import { defineMessages } from 'react-intl'
import { connect } from 'react-redux'
import { Container, Col, Row } from 'react-grid-system'
import { push } from 'connected-react-router'
import { bindActionCreators } from 'redux'

import Form from '@ttn-lw/components/form'
import PageTitle from '@ttn-lw/components/page-title'
import SubmitButton from '@ttn-lw/components/submit-button'
import Input from '@ttn-lw/components/input'
import SubmitBar from '@ttn-lw/components/submit-bar'

import OwnersSelect from '@console/containers/owners-select'

import getHostFromUrl from '@ttn-lw/lib/host-from-url'
import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'
import tooltipIds from '@ttn-lw/lib/constants/tooltip-ids'
import Yup from '@ttn-lw/lib/yup'
import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'
import { selectGsConfig } from '@ttn-lw/lib/selectors/env'
import { url as urlRegexp } from '@ttn-lw/lib/regexp'

import { id as idRegexp, address as addressRegexp } from '@console/lib/regexp'

import { claimGateway } from '@console/store/actions/gateways'

import { selectUserId } from '@console/store/selectors/user'

const m = defineMessages({
  cupsUri: 'CUPS Uri',
})

const validationSchema = Yup.object({
  authenticated_identifiers: Yup.object({
    gateway_eui: Yup.nullableString().length(8 * 2, Yup.passValues(sharedMessages.validateLength)),
    authentication_code: Yup.string(),
  }),
  collaborator: Yup.string()
    .min(2, Yup.passValues(sharedMessages.validateTooShort))
    .max(25, Yup.passValues(sharedMessages.validateTooLong))
    .matches(idRegexp, Yup.passValues(sharedMessages.validateIdFormat))
    .required(sharedMessages.validateRequired),
  target_gateway_id: Yup.string()
    .matches(idRegexp, Yup.passValues(sharedMessages.validateIdFormat))
    .min(2, Yup.passValues(sharedMessages.validateTooShort))
    .max(36, Yup.passValues(sharedMessages.validateTooLong))
    .required(sharedMessages.validateRequired),
  target_gateway_server: Yup.string()
    .matches(addressRegexp, Yup.passValues(sharedMessages.validateAddressFormat))
    .required(sharedMessages.validateRequired),
  target_cups_uri: Yup.string()
    .matches(urlRegexp, Yup.passValues(sharedMessages.validateUrl))
    .required(sharedMessages.validateRequired),
})

const defaultValues = {
  authenticated_identifiers: {
    gateway_eui: '',
    authentication_code: '',
  },
  collaborator: '',
  target_gateway_id: '',
  target_gateway_server: '',
  target_cups_uri: '',
}

const GatewayClaim = props => {
  const { claimGateway, onClaimSuccess, gsConfig, userId } = props

  const { enabled: gsEnabled, base_url: gsUrl } = gsConfig
  const initialValues = React.useMemo(
    () => ({
      ...defaultValues,
      collaborator: userId,
      target_gateway_server: gsEnabled ? getHostFromUrl(gsUrl) : '',
    }),
    [gsEnabled, gsUrl, userId],
  )

  const [error, setError] = React.useState('')
  const handleSubmit = React.useCallback(
    async values => {
      try {
        await claimGateway(values)

        const { target_gateway_id } = values
        onClaimSuccess(target_gateway_id)
      } catch (error) {
        setError(error)
      }
    },
    [claimGateway, onClaimSuccess],
  )

  return (
    <Container>
      <PageTitle tall title={sharedMessages.claimGateway} />
      <Row>
        <Col md={10} lg={9}>
          <Form
            error={error}
            onSubmit={handleSubmit}
            initialValues={initialValues}
            validationSchema={validationSchema}
          >
            <Form.SubTitle title={sharedMessages.claimGateway} />
            <OwnersSelect name="collaborator" required autoFocus />
            <Form.Field
              required
              title={sharedMessages.gatewayID}
              name="target_gateway_id"
              placeholder={sharedMessages.gatewayIdPlaceholder}
              component={Input}
              tooltipId={tooltipIds.GATEWAY_ID}
            />
            <Form.Field
              required
              title={sharedMessages.gatewayEUI}
              name="authenticated_identifiers.gateway_eui"
              type="byte"
              min={8}
              max={8}
              placeholder={sharedMessages.gatewayEUI}
              component={Input}
              tooltipId={tooltipIds.GATEWAY_EUI}
            />
            <Form.Field
              required
              title={sharedMessages.claimAuthCode}
              name="authenticated_identifiers.authentication_code"
              component={Input}
            />
            <Form.Field
              required
              title={sharedMessages.gatewayServerAddress}
              description={sharedMessages.gsServerAddressDescription}
              placeholder={sharedMessages.addressPlaceholder}
              name="target_gateway_server"
              component={Input}
            />
            <Form.Field
              required
              title={m.cupsUri}
              placeholder="https://example.thethings.com:443"
              name="target_cups_uri"
              component={Input}
            />
            <SubmitBar>
              <Form.Submit component={SubmitButton} message={sharedMessages.saveChanges} />
            </SubmitBar>
          </Form>
        </Col>
      </Row>
    </Container>
  )
}

GatewayClaim.propTypes = {
  claimGateway: PropTypes.func.isRequired,
  gsConfig: PropTypes.stackComponent.isRequired,
  onClaimSuccess: PropTypes.func.isRequired,
  userId: PropTypes.string.isRequired,
}

export default connect(
  state => ({
    gsConfig: selectGsConfig(),
    userId: selectUserId(state),
  }),
  dispatch => ({
    ...bindActionCreators({
      claimGateway: attachPromise(claimGateway),
    }),
    onClaimSuccess: gtwId => dispatch(push(`/gateways/${gtwId}`)),
  }),
)(GatewayClaim)
