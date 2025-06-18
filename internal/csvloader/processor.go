package csvloader

type CSVProcessor interface {
	Process() error
}

type processorWrapper[R any, D interface{ Append(*R) }] struct {
	*Processor[R, D]
}

func (pw *processorWrapper[R, D]) Process() error {
	return ProcessCSV(pw.Processor)
}

func WrapProcessor[R any, D interface{ Append(*R) }](p *Processor[R, D]) CSVProcessor {
	return &processorWrapper[R, D]{p}
}

type ParseFunc[R any] func([]string) (*R, error)

type Processor[R any, D interface{ Append(*R) }] struct {
	fileName  string
	driver    D
	parseFunc ParseFunc[R]
}

func NewProcessor[R any, D interface{ Append(*R) }](filename string, driver D, parseFunc ParseFunc[R]) *Processor[R, D] {
	return &Processor[R, D]{
		fileName:  filename,
		driver:    driver,
		parseFunc: parseFunc,
	}
}

func (p *Processor[R, D]) processCSVLine(line []string) error {
	obj, err := p.parseFunc(line)
	if err != nil {
		return err
	}
	p.driver.Append(obj)
	return nil
}

func (p *Processor[R, D]) Filename() string {
	return p.fileName
}
