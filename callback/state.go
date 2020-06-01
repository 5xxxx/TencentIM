/*
 *
 * state.go
 * callback
 *
 * Created by lintao on 2020/6/1 5:16 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package callback

type Status interface {
	StatueChange(contentType string, clientIP string, optPlatform string) error
}
