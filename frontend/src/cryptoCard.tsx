import { Pin, PinOff } from "lucide-react";

interface CurrencyPrices {
  [currency: string]: number;
}


interface CryptoCardProps {
  crypto: string;
  prices: CurrencyPrices;
  onPin: (crypto: string) => void;
  isPinned: boolean;
}

export default function CryptoCard({ crypto, prices, onPin, isPinned }: CryptoCardProps) {
  return (
    <div className="crypto-card">
      <div className="card-header mb-2">
        <h2 className="card-title text-xl font-semibold">{crypto}</h2>
        <button
          onClick={() => onPin(crypto)}
          className="card-pin p-0.5 rounded hover:bg-gray-200"
          title={isPinned ? "Unpin" : "Pin"}
          aria-label={isPinned ? "Unpin" : "Pin"}
        >
          {isPinned ? (
            <PinOff size={14} className="text-red-500" />
          ) : (
            <Pin size={14} className="text-gray-600" />
          )}
        </button>
      </div>

      {Object.keys(prices).map((currency) => (
        <div key={currency} className="text-lg">
          {currency} Price: ${prices[currency]?.toFixed(2)}
        </div>
      ))}
    </div>
  );
}
