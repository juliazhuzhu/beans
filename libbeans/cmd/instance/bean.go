package instance

import(
	"xiaoye.com/dory/beans/libbeans/bean"
	"context"
	svc "xiaoye.com/dory/beans/libbeans/service"
)
// Beat provides the runnable and configurable instance of a beat.
type Beat struct {
	bean.Bean

	name string
}

func Run(bt bean.Creator){

	b := bean.Bean{}
	bb := &Beat{Bean: b,
			name:"instance",}
	bb.launch(bt)
}

func(b* Beat)launch(bt bean.Creator) error{


	_, cancel := context.WithCancel(context.Background())
	
	bean, _ := bt(&b.Bean)
	svc.HandleSignals(bean.Stop, cancel)
	return bean.Run(&b.Bean)
	
}