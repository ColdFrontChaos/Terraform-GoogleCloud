
�
google/protobuf/empty.protogoogle.protobuf"
EmptyB}
com.google.protobufB
EmptyProtoPZ.google.golang.org/protobuf/types/known/emptypb��GPB�Google.Protobuf.WellKnownTypesJ�
 3
�
 2� Protocol Buffers - Google's data interchange format
 Copyright 2008 Google Inc.  All rights reserved.
 https://developers.google.com/protocol-buffers/

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are
 met:

     * Redistributions of source code must retain the above copyright
 notice, this list of conditions and the following disclaimer.
     * Redistributions in binary form must reproduce the above
 copyright notice, this list of conditions and the following disclaimer
 in the documentation and/or other materials provided with the
 distribution.
     * Neither the name of Google Inc. nor the names of its
 contributors may be used to endorse or promote products derived from
 this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.


  

" ;
	
%" ;

# E
	
# E

$ ,
	
$ ,

% +
	
% +

& "
	

& "

' !
	
$' !

( 
	
( 
�
 3 � A generic empty message that you can re-use to avoid defining duplicated
 empty messages in your APIs. A typical example is to use it as the request
 or the response type of an API method. For instance:

     service Foo {
       rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);
     }

 The JSON representation for `Empty` is empty JSON object `{}`.



 3bproto3
�+
bookstore.protoendpoints.examples.bookstoregoogle/protobuf/empty.proto"-
Shelf
id (Rid
theme (	Rtheme"D
Book
id (Rid
author (	Rauthor
title (	Rtitle"T
ListShelvesResponse=
shelves (2#.endpoints.examples.bookstore.ShelfRshelves"O
CreateShelfRequest9
shelf (2#.endpoints.examples.bookstore.ShelfRshelf"'
GetShelfRequest
shelf (Rshelf"*
DeleteShelfRequest
shelf (Rshelf"(
ListBooksRequest
shelf (Rshelf"M
ListBooksResponse8
books (2".endpoints.examples.bookstore.BookRbooks"a
CreateBookRequest
shelf (Rshelf6
book (2".endpoints.examples.bookstore.BookRbook":
GetBookRequest
shelf (Rshelf
book (Rbook"=
DeleteBookRequest
shelf (Rshelf
book (Rbook2�
	BookstoreZ
ListShelves.google.protobuf.Empty1.endpoints.examples.bookstore.ListShelvesResponse" f
CreateShelf0.endpoints.examples.bookstore.CreateShelfRequest#.endpoints.examples.bookstore.Shelf" `
GetShelf-.endpoints.examples.bookstore.GetShelfRequest#.endpoints.examples.bookstore.Shelf" Y
DeleteShelf0.endpoints.examples.bookstore.DeleteShelfRequest.google.protobuf.Empty" n
	ListBooks..endpoints.examples.bookstore.ListBooksRequest/.endpoints.examples.bookstore.ListBooksResponse" c

CreateBook/.endpoints.examples.bookstore.CreateBookRequest".endpoints.examples.bookstore.Book" ]
GetBook,.endpoints.examples.bookstore.GetBookRequest".endpoints.examples.bookstore.Book" W

DeleteBook/.endpoints.examples.bookstore.DeleteBookRequest.google.protobuf.Empty" B;
'com.google.endpoints.examples.bookstoreBBookstoreProtoPJ�
 }
�
 2� Copyright 2016 Google Inc. All Rights Reserved.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.

//////////////////////////////////////////////////////////////////////////////


 %

 "
	

 "

 /
	
 /

 @
	
 @
	
  %
k
  /_ A simple Bookstore API.

 The API manages shelves and books resources. Shelves contain books.



 
>
   I1 Returns a list of all shelves in the bookstore.


   

   '

   2E
4
 "8' Creates a new shelf in the bookstore.


 "

 "$

 "/4
2
 $2% Returns a specific bookstore shelf.


 $

 $

 $).
Q
 &HD Deletes a shelf, including all books that are stored on the shelf.


 &

 &$

 &/D
2
 (@% Returns a list of books on a shelf.


 (

 ( 

 (+<
"
 *5 Creates a new book.


 *

 *"

 *-1
'
 ,/ Returns a specific book.


 ,

 ,

 ,'+
+
 .F Deletes a book from a shelf.


 .

 ."

 .-B

 2 7 A shelf resource.



 2
!
  4 A unique shelf id.


  4

  4


  4
;
 6. A theme of the shelf (fiction, poetry, etc).


 6

 6	

 6

: A A book resource.



:
 
 < A unique book id.


 <

 <


 <
%
> An author of the book.


>

>	

>

@ A book title.


@

@	

@
+
D G Response to ListShelves call.



D
(
 F Shelves in the bookstore.


 F


 F

 F

 F
5
J M) Request message for CreateShelf method.



J
,
 L The shelf resource to create.


 L

 L

 L
2
P S& Request message for GetShelf method.



P
8
 R+ The ID of the shelf resource to retrieve.


 R

 R

 R
5
V Y) Request message for DeleteShelf method.



V
-
 X  The ID of the shelf to delete.


 X

 X

 X
3
\ _' Request message for ListBooks method.



\
3
 ^& ID of the shelf which books to list.


 ^

 ^

 ^
3
b e' Response message to ListBooks method.



b
&
 d The books on the shelf.


 d


 d

 d

 d
4
h m( Request message for CreateBook method.



h
=
 j0 The ID of the shelf on which to create a book.


 j

 j

 j
6
l) A book resource to create on the shelf.


l

l

l
1
	p u% Request message for GetBook method.



	p
A
	 r4 The ID of the shelf from which to retrieve a book.


	 r

	 r

	 r
.
	t! The ID of the book to retrieve.


	t

	t

	t
4

x }( Request message for DeleteBook method.




x
?

 z2 The ID of the shelf from which to delete a book.



 z


 z


 z
,

| The ID of the book to delete.



|


|


|bproto3