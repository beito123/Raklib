package raklib

/*
	Raklib

	Copyright (c) 2018 beito

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
*/

type NoSetBufferError struct {
}

func (e NoSetBufferError) Error() string {
	return "The buffer is't set."
}