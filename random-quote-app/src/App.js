import React, {useEffect, useState} from 'react';
import './App.css';

function App() {
  const [time, setTime] = useState(new Date());
  const [quote, setQuote] = useState('');
  const [author, setAuthor] = useState('');
  const [category, setCategory] = useState('');
  const [price, setBitcoinPrice] = useState(0.0);

  function capitalizeFirstLetter(string) {
    return string.charAt(0).toUpperCase() + string.slice(1);
  }
  const fetchRandomQuote = async () => {
    try {
      const response = await fetch('https://go-backend.tomgur.me:8080/random-quote');
      const data = await response.json();
      console.info("quote: " + data.QUOTE)
      console.info("author: " + data.AUTHOR)
      console.info("category: " + data.CATEGORY)
      var fixedCategory = capitalizeFirstLetter(data.CATEGORY)
      setQuote(data.QUOTE);
      setAuthor(data.AUTHOR);
      setCategory(fixedCategory);
    } catch (error) {
      console.error('Error fetching quote:', error);
    }
  };

  const fetchBitcoinPrice = async () => {
    try {
      const response = await fetch('https://go-backend.tomgur.me:8080/bitcoin-price');
      const data = await response.json();
      console.info("price: " + data)
      setBitcoinPrice("$" + data);
    } catch (error) {
      console.error('Error fetching price:', error);
    }
  }

  const handleQuoteChange = (event) => {
    setQuote(event.target.value);
    setAuthor(event.target.value);
    setCategory(event.target.value);
  };
  useEffect(() => {
    const interval = setInterval(() => {
      setTime(new Date());
    }, 1000);

    return () => {
      clearInterval(interval);
    };
  }, []);
const handlePriceChange = (event) => {
    setBitcoinPrice(event.target.value);
};

  return (
      <div className="App" style={{ backgroundColor: 'lightblue', minHeight: '100vh', display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
        <div className="centered">
          <p>{time.toLocaleTimeString()}</p>
          <textarea
              placeholder="Random quote will appear here"
              value={quote}
              cols={30}
              rows={10}
              readOnly
              onChange={handleQuoteChange}
          /><br></br>
          Author: <input
              type="text"
              placeholder="Author"
              value={author||''}
              onChange={handleQuoteChange}
          /><br></br>
          Category: <input
              type="text"
              placeholder="Category"
              value={category||''}
              onChange={handleQuoteChange}
          /><br></br>
          <button onClick={fetchRandomQuote}>Get Random Quote</button>
          <p>Bitcoin Price</p>
          <input
              type="text"
              placeholder="Bitcoin Price"
              value={price||''}
              onChange={handlePriceChange}
          />
          <button onClick={fetchBitcoinPrice}>Get Bitcoin Price</button>
        </div>
      </div>
  );
}

export default App;
