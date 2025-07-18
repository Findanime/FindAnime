import { Search } from 'lucide-react';

interface SearchBarProps {
  inputClassName?: string;
  buttonClassName?: string;
  iconSize?: number;
  containerClassName?: string;
  positionStyle?: React.CSSProperties;

  inputStyle?: React.CSSProperties;    // Direct <input> styling
  buttonStyle?: React.CSSProperties;   // Optional <button> styling

  width?: string | number;             // Bar width (e.g., '100%', '500px')
  height?: string | number;            // Bar height (e.g., '48px')
}

export default function SearchBar({
  inputClassName = '',
  buttonClassName = '',
  iconSize = 20,
  containerClassName = '',
  positionStyle = {},
  inputStyle = {},
  buttonStyle = {},
  width,
  height,
}: SearchBarProps) {
  return (
    <div className={`relative ${containerClassName}`} style={positionStyle}>
      <input
        type="text"
        placeholder="Search anime..."
        className={`rounded-full border border-white/20 bg-white/5 pl-6 pr-14 text-lg text-white placeholder-white/50 backdrop-blur-sm focus:border-pink-500 focus:outline-none focus:ring-2 focus:ring-pink-500/50 ${inputClassName}`}
        style={{
          width: width || '100%',
          height: height || 'auto',
          paddingTop: height ? undefined : '0.75rem',
          paddingBottom: height ? undefined : '0.75rem',
          ...inputStyle,
        }}
      />
      <button
        className={`absolute right-2 top-1/2 -translate-y-1/2 rounded-full bg-pink-600 p-2.5 text-white transition-colors hover:bg-pink-700 ${buttonClassName}`}
        style={buttonStyle}
      >
        <Search size={iconSize} />
      </button>
    </div>
  );
}
