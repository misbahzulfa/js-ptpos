package exception

import "errors"

/*---

Error Code

JMGE001 = General Error/ Error yang bersifat umum
JMGE002 = AWS Exception
JMGE003 = Environment Exception
JMGE004 = Call Api Exception
JMGE005 = JSON Parse Exception
JMGE006 = Client Do Exception
JMGE007 = Redis Exception
JMGE008 = JSON Parse From Redis Exception

*/

var (
	GeneralExceptionMessage        = errors.New("Terjadi kesalahan, silahkan ulangi beberapa saat lagi, code [JMGE001]")
	AWSExceptionMessage            = errors.New("Terjadi kesalahan, silahkan ulangi beberapa saat lagi, code [JMGE002]")
	EnvExceptionMessage            = errors.New("Terjadi kesalahan, silahkan ulangi beberapa saat lagi, code [JMGE003]")
	CallApiExceptionMessage        = errors.New("Terjadi kesalahan, silahkan ulangi beberapa saat lagi, code [JMGE004]")
	JSONParseExceptionMessage      = errors.New("Terjadi kesalahan, silahkan ulangi beberapa saat lagi, code [JMGE005]")
	ClientDoExceptionMessage       = errors.New("Terjadi kesalahan, silahkan ulangi beberapa saat lagi, code [JMGE006]")
	RedisExceptionMessage          = errors.New("Terjadi kesalahan, silahkan ulangi beberapa saat lagi, code [JMGE007]")
	TokenExceptionMessage          = errors.New("Sessi anda telah habis")
	JSONParseRedisExceptionMessage = errors.New("Terjadi kesalahan, silahkan ulangi beberapa saat lagi, code [JMGE008]")
)
