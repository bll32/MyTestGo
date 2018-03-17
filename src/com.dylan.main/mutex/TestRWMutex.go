package main

import "sync"

//只分析
/*RWMutex是一个读写锁，该锁可以加多个读锁或者一个写锁，其经常用于读次数远远多于写次数的场景．

func (rw *RWMutex) Lock()　　写锁，如果在添加写锁之前已经有其他的读锁和写锁，则lock就会阻塞直到该锁可用，为确保该锁最终可用，已阻塞的 Lock 调用会从获得的锁中排除新的读取器，即写锁权限高于读锁，有写锁时优先进行写锁定
func (rw *RWMutex) Unlock()　写锁解锁，如果没有进行写锁定，则就会引起一个运行时错误．
func (rw *RWMutex) RLock() 读锁，当有写锁时，无法加载读锁，当只有读锁或者没有锁时，可以加载读锁，读锁可以加载多个，所以适用于＂读多写少＂的场景
//func (rw *RWMutex)RUnlock()　读锁解锁，RUnlock 撤销单次 RLock 调用，它对于其它同时存在的读取器则没有效果。若 rw 并没有为读取而锁定，调用 RUnlock 就会引发一个运行时错误(注：这种说法在go1.3版本中是不对的，例如下面这个例子)。
*/
// 读锁可以多次加入


func main() {
var l *sync.RWMutex
l= new(sync.RWMutex)
//l.Lock()
l.RLock()
l.RLock()
l.RLock()
l.RLock()

}
