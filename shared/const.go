package shared

const SUCCESS = "success"
const MESSAGE = "message"
const DATA = "data"
const TOTAL = "total"
const ERROR = "error"

const InsufficientParamErr = "insufficient parameters, please refresh the page and try again"
const BadParamErr = "incompatible parameters, please refresh the page and try again"
const BadIdErr = "wrong format of ID provided"
const NoInfoErr = "no content found"
const InternalServerErr = "internal error occurred"
const UserNotFoundErr = "user info does not exist"
const ResourceAlreadyExistErr = "data already exist, please try another one"
const InfoMismatchErr = "input data doesn't match"
const NoPreviousRecordExistErr = "user has not provided needed information yet"

const MAX = 10e10

const CtxUserIDKey = "user_id"
const CtxUserFirstNameKey = "first_name"
const CtxUserLastNameKey = "last_name"
const CtxUserRole = "role"
