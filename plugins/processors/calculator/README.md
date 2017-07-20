# Calculator Processor Plugin

The calculator processor plugin performs four arithmetic operations between two fields and then create new field.

* add
* sub
* mul
* div

### Configuration:

```toml
# Print all metrics that pass through this filter.
[[processors.calculator]]
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
```
