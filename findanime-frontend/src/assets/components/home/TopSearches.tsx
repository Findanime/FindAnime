interface Props {
  searches: string[];
  className?: string;
  style?: React.CSSProperties;
  label?: string;
}

export default function TopSearches({
  searches,
  className = '',
  style,
  label = 'Top search:',
}: Props) {
  return (
    <div
      className={`text-sm text-white/80 py-0 ${className}`} // 👈 vertical padding set to 0
      style={{ maxWidth: '700px', paddingTop: 0, paddingBottom: 0, ...style }} // 👈 inline padding override
    >
      <div className="flex flex-wrap items-center gap-x-3 gap-y-2">
        <span className="font-bold text-white whitespace-nowrap">{label}</span>
        {searches.map((item, index) => (
          <a
            key={index}
            href="#"
            className="text-white/70 hover:text-pink-400 transition-colors"
          >
            {item}
          </a>
        ))}
      </div>
    </div>
  );
}
