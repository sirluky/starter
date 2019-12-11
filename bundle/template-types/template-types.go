package template_types

type TemplateTypes struct {
	Stencils        []*Stencil        `json:"stencils"`
	Policies        []*Policy         `json:"policies"`
	Transformations []*Transformation `json:"transformations"`
	HelmCharts      []*HelmRelease    `json:"helm_charts"`
	Workflows       []*Workflow       `json:"workflows"`
}
