package fmp4Tomp4

// func Test_w(t *testing.T) {
// 	tf := file.New("1.mp4", 0, false)
// 	if tf.IsExist() {
// 		_ = tf.Delete()
// 	}
// 	defer tf.Close()

// 	w, e := NewBoxWriter(tf)
// 	if e != nil {
// 		t.Fatal(e)
// 	}

// 	// ftyp
// 	{
// 		ftyp := w.Box("ftyp")
// 		ftyp.Write([]byte("isom"))
// 		ftyp.Write(itob32(512))
// 		ftyp.Write([]byte("isom"))
// 		ftyp.Write([]byte("iso2"))
// 		ftyp.Write([]byte("avc1"))
// 		ftyp.Write([]byte("mp41"))
// 		if e := ftyp.Close(); e != nil {
// 			t.Fatal(e)
// 		}
// 	}
// }
