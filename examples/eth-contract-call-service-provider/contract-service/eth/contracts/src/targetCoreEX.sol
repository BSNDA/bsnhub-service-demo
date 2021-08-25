//SPDX-License-Identifier: SimPL-2.0
pragma solidity ^0.8.7;

/**
 * @title iService Core Extension contract
 */
contract targetCoreEx {
    mapping(bytes32 => bool) requests;

    /**
    * @dev Event triggered when the request is sent
    * @param _RequestID Request id
    * @param _result result bytes
    */
    event CrossChainResponseSent(
        bytes32 _RequestID,
        bytes _result
    );
    /**
    * @dev Make sure that the Request  has not been responded
    * @param _RequestID Request id
    */
    modifier validateRequest(bytes32 _RequestID) {
        require(
            requests[_RequestID] == false,
            "iServiceCoreEx: duplicated request!"
        );

        _;
    }

    /**
     * @dev call service/contract in dest chain
     * @param _RequestID Request id
     * @param _endpointAddress endpointAddress
     * @param _callData call data from source chain
     */
    function callService(
        bytes32 _RequestID,
        address _endpointAddress,
        bytes memory _callData
    ) public validateRequest(_RequestID) {
        (bool success, bytes memory result) = _endpointAddress.call(_callData);
        if (success == true) {
            emit CrossChainResponseSent(
                _RequestID,
                result
            );
            requests[_RequestID] = true;
        }
    }
}
