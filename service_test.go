package psqlsrv_test

import (
	"github.com/lab259/go-rscsrv"
	. "github.com/lab259/go-rscsrv-psql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Psql Service", func() {
	It("should fail loading a configuration", func() {
		var service PsqlService
		configuration, err := service.LoadConfiguration()
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(ContainSubstring("not implemented"))
		Expect(configuration).To(BeNil())
	})

	It("should fail applying configuration", func() {
		var service PsqlService
		err := service.ApplyConfiguration(map[string]interface{}{
			"address": "localhost",
		})
		Expect(err).To(Equal(rscsrv.ErrWrongConfigurationInformed))
	})

	It("should apply the configuration using a pointer", func() {
		var service PsqlService
		err := service.ApplyConfiguration(&Configuration{
			Host:        "host",
			Username:    "user",
			Password:    "password",
			Database:    "database",
			Port:        3306,
			MaxPoolSize: 1,
		})
		Expect(err).To(BeNil())
		Expect(service.Configuration.Host).To(Equal("host"))
		Expect(service.Configuration.Username).To(Equal("user"))
		Expect(service.Configuration.Password).To(Equal("password"))
		Expect(service.Configuration.Database).To(Equal("database"))
		Expect(service.Configuration.Port).To(Equal(3306))
		Expect(service.Configuration.MaxPoolSize).To(Equal(1))
	})

	It("should apply the configuration using a pointer", func() {
		var service PsqlService
		err := service.ApplyConfiguration(Configuration{
			Host:        "host",
			Username:    "user",
			Password:    "password",
			Database:    "database",
			Port:        3306,
			MaxPoolSize: 1,
		})
		Expect(err).To(BeNil())
		Expect(service.Configuration.Host).To(Equal("host"))
		Expect(service.Configuration.Username).To(Equal("user"))
		Expect(service.Configuration.Password).To(Equal("password"))
		Expect(service.Configuration.Database).To(Equal("database"))
		Expect(service.Configuration.Port).To(Equal(3306))
		Expect(service.Configuration.MaxPoolSize).To(Equal(1))
	})

	It("should start the service", func() {
		var service PsqlService
		Expect(service.ApplyConfiguration(Configuration{
			Username:    "postgres",
			Password:    "postgres",
			Database:    "postgres",
			Port:        5432,
			MaxPoolSize: 1,
		})).To(Succeed())
		Expect(service.Start()).To(Succeed())
		defer service.Stop()
	})

	It("should stop the service", func() {
		var service PsqlService
		Expect(service.ApplyConfiguration(Configuration{
			Username:    "postgres",
			Password:    "postgres",
			Database:    "postgres",
			Port:        5432,
			MaxPoolSize: 1,
		})).To(Succeed())
		Expect(service.Start()).To(Succeed())
		Expect(service.Stop()).To(Succeed())
		Expect(service.Ping()).To(Equal(rscsrv.ErrServiceNotRunning))
		_, err := service.DB()
		Expect(err).To(Equal(rscsrv.ErrServiceNotRunning))
	})

	It("should restart the service", func() {
		var service PsqlService
		Expect(service.ApplyConfiguration(Configuration{
			Username:    "postgres",
			Password:    "postgres",
			Database:    "postgres",
			Port:        5432,
			MaxPoolSize: 1,
		})).To(Succeed())
		Expect(service.Start()).To(Succeed())
		Expect(service.Restart()).To(Succeed())
		Expect(service.Ping()).To(Succeed())
	})
})
