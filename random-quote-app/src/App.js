import React, { useState, useEffect } from 'react';
import { googleLogout, useGoogleLogin } from '@react-oauth/google';
import axios from 'axios';


function App() {
    const [ user, setUser ] = useState([]);
    const [ profile, setProfile ] = useState([]);
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

    const login = useGoogleLogin({
        onSuccess: (codeResponse) => setUser(codeResponse),
        onError: (error) => console.log('Login Failed:', error)
    });

    useEffect(
        () => {
            if (user) {
                axios
                    .get(`https://www.googleapis.com/oauth2/v1/userinfo?access_token=${user.access_token}`, {
                        headers: {
                            Authorization: `Bearer ${user.access_token}`,
                            Accept: 'application/json'
                        }
                    })
                    .then((res) => {
                        setProfile(res.data);
                    })
                    .catch((err) => console.log(err));
            }
        },
        [ user ]
    );

    // log out function to log the user out of google and set the profile array to null
    const logOut = () => {
        googleLogout();
        setProfile(null);
    };

    return (
        <div>
            <h2>React Google Login</h2>
            <br />
            <br />
            {profile ? (
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
                    />
                    <br />
                    <br />
                    Category: <input
                        type="text"
                        placeholder="Category"
                        value={category||''}
                        onChange={handleQuoteChange}
                    />
                    <button onClick={fetchRandomQuote}>Get Random Quote</button>
                    <br />
                    <br />
                    <p>Bitcoin Price</p>
                    <input
                        type="text"
                        placeholder="Bitcoin Price"
                        value={price||''}
                        onChange={handlePriceChange}
                    />
                    <button onClick={fetchBitcoinPrice}>Get Bitcoin Price</button>
                    <br />
                    <br />
                    <img src={profile.picture} alt="user image" />
                    <h3>User Logged in</h3>
                    <p>Name: {profile.name}</p>
                    <p>Email Address: {profile.email}</p>
                    <br />
                    <br />
                    <button onClick={logOut}>Log out</button>

                </div>
            ) : (
                <button onClick={() => login()}>Sign in with Google 🚀 </button>
            )}
        </div>
    );
}
export default App;