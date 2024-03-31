package scenes

type SceneManager struct {
	scenes []Scene
}

var sceneManager SceneManager

func Init(scene Scene) {
	scenes := make([]Scene, 1, 5)
	scenes[0] = scene
	sceneManager = SceneManager{
		scenes: scenes,
	}
}

func GetManager() *SceneManager {
	return &sceneManager
}

func (manager *SceneManager) Push(scene Scene) {
	manager.scenes = append(manager.scenes, scene)
}

func (manager *SceneManager) GetCurrent() Scene {
	return manager.scenes[len(manager.scenes)-1]
}

func (manager *SceneManager) Pop() Scene {
	prev := manager.scenes[len(manager.scenes)-1]
	manager.scenes = manager.scenes[:len(manager.scenes)-1]
	return prev
}

func (manager *SceneManager) Swap(scene Scene) {
	manager.scenes[len(manager.scenes)-1] = scene
}
