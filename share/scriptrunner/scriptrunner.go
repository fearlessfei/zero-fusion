package scriptrunner

// ScriptLogicRunner 代表一个脚本业务
type ScriptLogicRunner interface {
	Run() error
}
