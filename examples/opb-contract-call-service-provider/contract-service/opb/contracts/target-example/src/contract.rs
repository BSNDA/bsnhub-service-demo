use cosmwasm_std::{
    to_binary, entry_point, attr, Deps, DepsMut, Env, Response, MessageInfo,Binary,StdResult
};

use crate::error::ContractError;
use crate::msg::{HandleMsg, InitMsg,QueryMsg};
use crate::state::{State,config};

// Note, you can use StdResult in some functions where you do not
// make use of the custom errors
#[cfg_attr(not(feature = "library"), entry_point)]
pub fn instantiate(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    _msg: InitMsg,
) -> Result<Response, ContractError> {

    let state = State {
        owner: deps.api.addr_canonicalize(&info.sender.to_string())?,
    };
    config(deps.storage).save(&state)?;

    Ok(Response::default())
}

// And declare a custom Error variant for the ones where you will want to make use of it
#[cfg_attr(not(feature = "library"), entry_point)]
pub fn execute(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    msg: HandleMsg,
) ->  Result<Response, ContractError> {
    match msg {
        HandleMsg::Hello{words} => try_hello(deps,words),
    }
}

pub fn try_hello(_deps: DepsMut, words: String) ->  Result<Response, ContractError> {
    let mut res = Response::default();
    res.data = Some(Binary::from(words.as_bytes()));
    res.attributes = vec![attr("result", words)];
    Ok(res)
}

pub fn query(_deps: Deps, _env: Env, _msg: QueryMsg) -> StdResult<Binary> {
    to_binary("no query function")
}
