package audio

// type mp3ReadCloser struct {
// 	decoder *mp3.Decoder
// 	file    *os.File
// }

// func newMP3ReadCloser(file *os.File) (*mp3ReadCloser, error) {
// 	decoder, err := mp3.NewDecoder(file)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &mp3ReadCloser{
// 		decoder: decoder,
// 		file:    file,
// 	}, nil
// }

// func (m *mp3ReadCloser) Read(p []byte) (int, error) {
// 	return m.decoder.Read(p)
// }

// func (m *mp3ReadCloser) Close() error {
// 	return m.file.Close()

// file, err := os.Open("./music/08 Piste 8.mp3")
// 	if err != nil {
// 		return
// 	}
// 	mp3Reader, err := newMP3ReadCloser(file)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer mp3Reader.Close()

// 	// Créer un nouveau contexte audio
// 	audioContext, err := audio.NewContext(44100)
// 	if err != nil {

// 	}
// 	player, err := audio.NewPlayer(audioContext, mp3Reader)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	player.SetVolume(0.8)
// 	player.Play()

// 	for player.IsPlaying() {
// 	}
