import React, { useEffect, useState } from 'react';

interface Book {
    Id: number;
    Name: string;
    Price: number;
}

const BookList: React.FC = () => {
    const [books, setBooks] = useState<Book[]>();

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch('http://localhost:8080');
                const data = await response.json();
                setBooks(data.books);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };

        fetchData();
    }, []);

    return (
        <div>
            <h1>Book List</h1>
            <ul>
                {books?.map((book) => (
                    <li key={book.Id}>
                        <strong>{book.Name}</strong> - ${book.Price}
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default BookList;
