package shake

type Snake struct {
	Parts []Point
}

func NewSnake() *Snake {
	snake := &Snake{}
	//...
	return snake
}

// определяем длину змейки(определяем длину массива)
func (s *Snake) Len() int {
	return len(s.Parts)
}

//len(), Head() - названия анонимных функций

// определяем по голове змейки где она находится
func (s *Snake) Head() Point {
	if s.Len() == 0 {
		return Point{-1, -1}
	}
	return s.Parts[0]
	//нулевая ячкйка массива - голова
}

// определяем по хвосту змейки где она находится
func (s *Snake) Tail() Point {
	if s.Len() == 0 {
		return Point{-1, -1}
	}
	return s.Parts[len(s.Parts)-1]
	//нулевая ячкйка массива - голова
}

// увеличение змейки
func (s *Snake) Add(p Point) {
	s.Parts = append([]Point{p}, s.Parts...)
}

//добавляем новый элемент вначало массива

// проверка если змейка натыкается на саму себя
func (s *Snake) IsSnake(p Point) bool {
	for i := range s.Parts {
		if s.Parts[i] == p {
			return true
		}
	}
	return false
}

func (s *Snake) CutIfSnake(p Point) bool {
	i := 0
	for ; i < len(s.Parts); i++ {
		if s.Parts[i] == p {
			break
		}
	}

	s.Parts = s.Parts[0:i] //обрезаем то что было до ячейки в котрую она врезалась

	return i >= len(s.Parts)
}

func (s *Snake) Reset() {
	//пишем начальные координаты для каждой новой игры
	sx, sy, l := 1, 1, 5
	for i := l - 1; i >= 0; i-- {
		s.Parts = append(s.Parts, Point{float64(sx + i), float64(sy)})
	}
}

func (s *Snake) Move(d Dir) {
	//надо подвинуть голову и по цепочке передвинуть остальные к голове
	lastP := s.Parts[0]
	s.Parts[0] = d.Exec(s.Parts[0])
	//в цикле передвинуть друг за другом все ячейки змейки кроме головы
	for i := range s.Parts[1:] {
		s.Parts[i+1], lastP = lastP, s.Parts[i+1]
		//координата меняется местами с предыдущей точкой

	}
}
