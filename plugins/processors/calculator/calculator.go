package calculator

import (
	"log"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/processors"
)

type Operation struct {
	FieldName string `toml:"field_name"`
	Operator  string `toml:"operator"`
	Left      string `toml:"left"`
	Right     string `toml:"right"`
}

type Calculator struct {
	Operations []Operation `toml:"operations"`
}

var sampleConfig = `
[[processors.calculator.operations]]
  field_name = "foo_total"
  operator = "mul"
  left = "foo_price"
  right = "foo_num"
[[processors.calculator.operations]]
  field_name = "bar_ave"
  operator = "div"
  left = "bar_total"
  right = "bar_num"
`

func (c *Calculator) SampleConfig() string {
	return sampleConfig
}

func (c *Calculator) Description() string {
	return "Perform four arithmetic operations between two fields."
}

func (c *Calculator) Apply(in ...telegraf.Metric) []telegraf.Metric {
	for _, metric := range in {
		fields := metric.Fields()
		// divisions
		for _, o := range c.Operations {
			if left, ok := fields[o.Left]; ok {
				log.Printf("D! Left: %v\n", left)
				//log.Printf("I! Left Type: %v\n", reflect.TypeOf(left))
				if left, ok := left.(int64); ok {
					log.Printf("D! Left: %v\n", left)
					if right, ok := fields[o.Right]; ok {
						log.Printf("D! Right: %v\n", right)
						if right, ok := right.(int64); ok {
							log.Printf("D! Right: %v\n", right)
							switch o.Operator {
							case "add":
								v := left + right
								log.Printf("D! %s: %v\n", o.FieldName, v)
								metric.AddField(o.FieldName, v)
							case "sub":
								v := left - right
								log.Printf("D! %s: %v\n", o.FieldName, v)
								metric.AddField(o.FieldName, v)
							case "mul":
								v := left * right
								log.Printf("D! %s: %v\n", o.FieldName, v)
								metric.AddField(o.FieldName, v)
							case "div":
								if right > 0 {
									v := left / right
									log.Printf("D! %s: %v\n", o.FieldName, v)
									metric.AddField(o.FieldName, v)
								}
							default:
								log.Printf("W! Unknown operator: %s\n", o.Operator)
							}
						}
					}
				}
			}
		}
	}
	return in
}

func init() {
	processors.Add("calculator", func() telegraf.Processor {
		return &Calculator{}
	})
}
