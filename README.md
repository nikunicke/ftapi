# ftAPI - WORK IN PROGRESS

WORK IN PROGRESS!
*   Bugs with float values. Some structs might have ints instead of floats, causing the json decoding to fail...
*   Needs tests
*   Nil values in ftapi.json response examples decodes to interface{} type, which isnt good. But can't be helped at this stage as the 42 API Docs don't provide accurate json schemas. If these fields would be used, they need to be type asserted first.

An API client for the 42 API.

