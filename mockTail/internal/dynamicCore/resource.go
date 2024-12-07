package dynamicCore

type SourceResource struct {
  Path      string
  Resources []Resource
}

type Resource struct {
  Method    string
  Name      string
  Token     []string
  HttpCode  int
  Params    map[string]string
  Extension string
  Error     *error
  IsValid   bool
}

func NewResource(parser Parser) (*Resource, error) {
  res := Resource{
    Name: parser.GetEntry().Name(),
  }

  if err := parser.Tokenize(&res); err != nil {
    return nil, err
  }
  if err := parser.ParseBody(&res); err != nil {
    return nil, err
  }
  if err := parser.ParseParams(&res); err != nil {
    res.Error = &err
  }
  if err := parser.ParseDynamic(&res); err != nil {
    return nil, err
  }
  if res.Error != nil {
    res.IsValid = false
  }
  return &res, nil
}
