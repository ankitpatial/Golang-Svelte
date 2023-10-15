/*
 * Copyright (c) 2022. SimSaw Software Private Limited.
 * Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 * Author: Ankit Patial
 * Dated:  05/04/22, 6:05 PM
 */

package open

// WithDefaultApp open a file, directory, or URI using the OS's default application for that object type.
func WithDefaultApp(input string) error {
	cmd := open(input)
	return cmd.Run()
}

//WithApp will open a file directory, or URI using the specified application.
//func WithApp(input string, appName string) error {
//	return openWith(input, appName).Run()
//}
