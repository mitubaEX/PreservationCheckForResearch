extract = function(type, from){
    parser = engine.parser()
    source = engine.source(from)
    obj = {}
    obj.time = sys.measure(function(){
        obj.birthmarks = parser.parse(source);
    })
    return obj;
}

compare = function(pairingAlgorithm, comparingAlgorithm, birthmarks1, birthmarks2){
    pair = engine.pairMatcher(pairingAlgorithm)
    comparator = engine.comparator(comparingAlgorithm)
    obj = {};
    obj.time = sys.measure(function(){
        obj.comparisons = comparator.compare(birthmarks1, birthmarks2, pair);
    });
    return obj;
}

extractResult1 = extract("2-gram", argv[1]);
extractResult2 = extract("2-gram", argv[2]);
compareResult = compare("RoundRobinWithSamePair", "JaccardIndex", extractResult1.birthmarks, extractResult2.birthmarks);
// compareResult = compare("RoundRobin", "JaccardIndex", extractResult.birthmarks);
// compareResult = compare("SameName", "JaccardIndex", extractResult.birthmarks);
// compareResult = compare("SameName", "JaccardIndex", extractResult1.birthmarks, extractResult2.birthmarks);

compareResult.comparisons.forEach(function(comparison){
    print(comparison);
});

// fs.print("extraction: " + extractResult.time + " ns")
fs.print("comparison: " + compareResult.time + " ns")


