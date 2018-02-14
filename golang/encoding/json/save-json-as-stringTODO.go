// I have a struct something like below, which has a field JsonStr that stores some json payload. 
// when the marshal happens, there should not be any json syntax errors because of the quotes in the json payload string. how to take care of this?
//
// in other words, the json payload string below should be like below in the marshal output
// payload example: // TODO add it here
// expected string in the marshal output under the JsonStr field: //TODO add it here, after escaping the doublequotes and slash characters.
//
