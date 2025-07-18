import { ArrowRight } from 'lucide-react';

interface WatchButtonProps {
  className?: string; // Tailwind classes for positioning, size, etc.
  style?: React.CSSProperties; // Inline positioning and sizing
  iconSize?: number; // Size of the ArrowRight icon
}

export default function WatchButton({
  className = '',
  style = {},
  iconSize = 24,
}: WatchButtonProps) {
  return (
    <button
      className={`group flex items-center gap-3 rounded-full bg-pink-600 px-8 py-4 text-lg font-bold text-white shadow-lg shadow-pink-600/30 transition-all hover:bg-pink-700 hover:shadow-xl hover:shadow-pink-600/40 ${className}`}
      style={style}
    >
      <span>Watch anime</span>
      <ArrowRight size={iconSize} className="transition-transform group-hover:translate-x-1" />
    </button>
  );
}
