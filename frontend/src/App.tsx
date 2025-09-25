import { useEffect, useState } from 'react'
import './App.css'
import CryptoCard from './cryptoCard'

function App() {
  const [selected, setSelected] = useState("")
  const [data, setData] = useState<CryptoData>({})
  const [pinned, setPinned] = useState<string[]>([]);


  const handlePin = (crypto: string) => {
    setPinned((prev) =>
      prev.includes(crypto) ? prev.filter((c) => c !== crypto) : [...prev, crypto]
    );
  };

  interface CurrencyPrices {
    [currency: string]: number;
  }

  interface CryptoData {
    [crypto: string]: CurrencyPrices;
  }

  const fetchCurrencyData = async () => {
    const res = await fetch(
      "http://localhost:9090/getAllCurrencies"
    );
    const respData = await res.json();
    console.log(respData)
    setData(respData["data"])
  }

  useEffect(() => {
    fetchCurrencyData()
  }, [])

  const handleChange = (e: any) => {
    setSelected(e.target.value);
  };

  return (
    
    <>
      {pinned.length > 0 && (
      <div className="pinned-section">
        <h3 className="pinned-title">Pinned Cryptos</h3>
        <div className="pinned-row">
          {pinned.map((crypto) => (
            <div key={crypto} className="pinned-item">
              <CryptoCard
                crypto={crypto}
                prices={data[crypto]}
                onPin={handlePin}
                isPinned={true}
              />
            </div>
          ))}
        </div>
      </div>
    )}

      <div className="selected-section">
        <select
          id="crypto"
          value={selected}
          onChange={handleChange}
          className="border rounded px-2 py-1"
        >
          <option value="" disabled>
            -- Choose --
          </option>
          {Object.keys(data).map((key) => (
            <option key={key} value={key}>
              {key}
            </option>
          ))}
        </select>
        {selected && data[selected] && (
          <CryptoCard
            crypto={selected}
            prices={data[selected]}
            onPin={handlePin}
            isPinned={pinned.includes(selected)}
          />
        )}
      </div>
    </>
  )
}

export default App
