// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package testtext contains test data that is of common use to the text
// repository.
package testtext // import "golang.org/x/text/internal/testtext"

const (

	// ASCII is an ASCII string containing all letters in the English alphabet.
	ASCII = "The quick brown fox jumps over the lazy dog. " +
		"The quick brown fox jumps over the lazy dog. " +
		"The quick brown fox jumps over the lazy dog. " +
		"The quick brown fox jumps over the lazy dog. " +
		"The quick brown fox jumps over the lazy dog. " +
		"The quick brown fox jumps over the lazy dog. " +
		"The quick brown fox jumps over the lazy dog. " +
		"The quick brown fox jumps over the lazy dog. " +
		"The quick brown fox jumps over the lazy dog. " +
		"The quick brown fox jumps over the lazy dog. "

	// Vietnamese is a snippet from http://creativecommons.org/licenses/by-sa/3.0/vn/
	Vietnamese = `Với các điều kiện sau: Ghi nhận công của tác giả. 
Nếu bạn sử dụng, chuyển đổi, hoặc xây dựng dự án từ 
nội dung được chia sẻ này, bạn phải áp dụng giấy phép này hoặc 
một giấy phép khác có các điều khoản tương tự như giấy phép này
cho dự án của bạn. Hiểu rằng: Miễn — Bất kỳ các điều kiện nào
trên đây cũng có thể được miễn bỏ nếu bạn được sự cho phép của
người sở hữu bản quyền. Phạm vi công chúng — Khi tác phẩm hoặc
bất kỳ chương nào của tác phẩm đã trong vùng dành cho công
chúng theo quy định của pháp luật thì tình trạng của nó không 
bị ảnh hưởng bởi giấy phép trong bất kỳ trường hợp nào.`

	// Russian is a snippet from http://creativecommons.org/licenses/by-sa/1.0/deed.ru
	Russian = `При обязательном соблюдении следующих условий:
Attribution — Вы должны атрибутировать произведение (указывать
автора и источник) в порядке, предусмотренном автором или
лицензиаром (но только так, чтобы никоим образом не подразумевалось,
что они поддерживают вас или использование вами данного произведения).
Υπό τις ακόλουθες προϋποθέσεις:`

	// Greek is a snippet from http://creativecommons.org/licenses/by-sa/3.0/gr/
	Greek = `Αναφορά Δημιουργού — Θα πρέπει να κάνετε την αναφορά στο έργο με τον
τρόπο που έχει οριστεί από το δημιουργό ή το χορηγούντο την άδεια
(χωρίς όμως να εννοείται με οποιονδήποτε τρόπο ότι εγκρίνουν εσάς ή
τη χρήση του έργου από εσάς). Παρόμοια Διανομή — Εάν αλλοιώσετε,
τροποποιήσετε ή δημιουργήσετε περαιτέρω βασ