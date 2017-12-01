// input -> birthmark, csv, csv

extractor = bmsys.extractor("2-gram");
source1 = fs.open(argv[1]);
source2 = fs.open(argv[2]);
p = extractor.extract(source1);
q = extractor.extract(source2);

// pair2 = bmsys.pairMaker("RoundRobinWithSamePair")
pair2 = bmsys.pairMaker("RoundRobin")
// comparator = bmsys.comparator("JaccardIndex")
// comparator = bmsys.comparator("DiceIndex")
comparator = bmsys.comparator("SimpsonIndex")

birthmarks = p.merge(q);

comparisons = comparator.compare(p, q, pair2);

fs.print(comparisons);
//
// var regExp = new RegExp("/", "g");
//
//
// var reg = /\//g;
// fs.writer("compare_result/" + argv[2].replace(reg,".") + "-" + argv[3].replace(reg,".") + ".csv", comparisons);

// fs.print("extraction: " + birthmarks.time() + " ns")
fs.print(comparisons.time() + " ns")
// fs.print("" + comparisons.time() + "")
