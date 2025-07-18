import SearchBar from './SearchBar';
import TopSearches from './TopSearches';
import WatchButton from './WatchButton';

const topSearches = [
  'Demon Slayer: Mt. Natagu...',
  'Dan Da Dan Season 2',
  'One Piece',
  'Kaiju No. 8 Season 2',
  'Demon Slayer: Kimetsu no ...',
  'Demon Slayer: The Hashira ...',
  'The Fragrant Flower Bloom...',
  'The Water Magician',
  'My Dress-Up Darling Seaso...',
  'Demon Slayer: Kimetsu no ...',
];

export default function SearchPane() {
  const boxStyles = {
    width: '1500px',           // ⬅️ customize width here
    height: '500px',           // ⬅️ customize height here
    marginTop: '150px'          // ⬅️ manually control vertical offset from top
  };

  return (
    <div className="w-full text-white font-sans antialiased flex justify-center px-4">
      <div
        className="rounded-2xl bg-[#1A181F00] shadow-lg shadow-white/2 backdrop-blur-sm p-4 flex flex-col justify-between"
        style={boxStyles}
      >
        <div className="absolute left-19 top-10">
          <h1 className="text-5xl font-bold text-white mb-2">
            F<span className="text-pink-500">!</span>ndAnime
          </h1>
        </div>
        <SearchBar
          width="700px"
          height="50px"
          iconSize={20}
          inputClassName="text-base"
          containerClassName="mt-6"
          positionStyle={{ marginLeft: '50px', marginRight: 'auto', marginTop: '120px', marginBottom: '0px' }}
        />
        <TopSearches searches={topSearches} style={{ fontSize: '14px', marginLeft: '60px', marginTop: "-100px" }} />
        <div className="flex justify-center">
          <WatchButton
            iconSize={30}
            className="absolute"
            style={{ top: '400px', left: '70px', width: '220px', height: '60px' }}
          />
        </div>
      </div>
    </div>
  );
}
