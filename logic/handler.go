package logic

type RequestHandler func(r RequestContextInterface)

type RequestContextProcessor func(ctx RequestContextInterface) RequestContextInterface
