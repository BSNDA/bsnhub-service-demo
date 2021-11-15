use cosmwasm_std::{
    to_binary, entry_point, Deps, Binary, HumanAddr, DepsMut, Env, Response, MessageInfo, WasmMsg, StdResult, SubMsg};

use crate::error::ContractError;
use crate::msg::{HandleMsg, InitMsg, QueryMsg};
use crate::state::REQUESTS;

// Note, you can use StdResult in some functions where you do not
// make use of the custom errors
#[cfg_attr(not(feature = "library"), entry_point)]
pub fn instantiate(
    _deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    _msg: InitMsg,
) -> Result<Response, ContractError> {
    Ok(Response::default())
}

// And declare a custom Error variant for the ones where you will want to make use of it
#[cfg_attr(not(feature = "library"), entry_point)]
pub fn execute(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    msg: HandleMsg,
) -> Result<Response, ContractError>  {
    match msg {
        HandleMsg::CallService{request_id, endpoint_address, call_data} 
        => call_service(deps, request_id, endpoint_address, call_data),
    }
}

pub fn call_service(deps: DepsMut, request_id:String, endpoint_address: HumanAddr, call_data: Binary) -> Result<Response, ContractError>  {
    let executed = REQUESTS.load(deps.storage, &request_id);
    if executed.is_ok() {
        Err(ContractError::Unauthorized{})
    }else{
        let single_msg = call_data;

        let messages =vec![SubMsg::new(WasmMsg::Execute {
            contract_addr: endpoint_address.to_string(),
            msg: single_msg,
            funds: vec![],
        })];

        REQUESTS.save(deps.storage, &request_id, &true)?;
        let mut res = Response::default();
        res.messages = messages;
        Ok(res)
    }
}

pub fn query(_deps: Deps, _env: Env, _msg: QueryMsg) -> StdResult<Binary> {
    to_binary("no query function")
}