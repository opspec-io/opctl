package models

func NewOpView(
description string,
name string,
subOps []OpSummaryView,
) *OpView {

  return &OpView{
    Description:description,
    Name:name,
    SubOps:subOps,
  }

}

type OpView struct {
  Description string
  Name        string
  SubOps      []OpSummaryView
}
