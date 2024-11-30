package math_test

import (
	"github.com/moleculer-go/moleculer/broker"
	test "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/moleculer-go/moleculer/examples/math"
)

var _ = test.Describe("Math", func() {

	test.It("Can create a valid service definition", func() {
		serviceDefinition := math.ServiceSchema()

		Expect(serviceDefinition).Should(Not(BeNil()))
		Expect(serviceDefinition.Name).To(Equal("math"))

		Expect(serviceDefinition.Actions).Should(HaveLen(3))
		Expect(serviceDefinition.Actions[0].Name).To(Equal("add"))
		Expect(serviceDefinition.Actions[1].Name).To(Equal("sub"))
		Expect(serviceDefinition.Actions[2].Name).To(Equal("mult"))

		Expect(serviceDefinition.Events).Should(HaveLen(2))
		Expect(serviceDefinition.Events[0].Name).To(Equal("math.add.called"))
		Expect(serviceDefinition.Events[1].Name).To(Equal("math.sub.called"))

	})

	test.It("Can start broker with service and call actions", func() {
		broker := broker.New()
		broker.Publish(math.ServiceSchema())
		broker.Start()

		Expect(broker).Should(Not(BeNil()))

		result := <-broker.Call("math.add", map[string]int{
			"a": 1,
			"b": 10,
		})

		Expect(result.Value()).Should(Not(BeNil()))
		Expect(result.Value()).Should(Equal(11))
	})
})
