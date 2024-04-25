export default function getInput() {
    const start = document.getElementById("startInput").value;
    const dest = document.getElementById("destInput").value;
    
    // Get the value of the switch
    const switchValue = document.querySelector('.switch input[type="checkbox"]').checked;
    
    // Clear input fields
    document.getElementById("startInput").value = "";
    document.getElementById("destInput").value = "";

    // Determine the search algorithm based on the switch value
    const algorithm = switchValue ? "IDS" : "BFS";

    // Use the collected data and chosen algorithm as needed
    console.log("Start:", start);
    console.log("Destination:", dest);
    console.log("Algorithm:", algorithm);

    // Perform further actions based on the chosen algorithm
    if (algorithm === "IDS") {
        // Perform actions for IDS algorithm
        console.log("Using IDS algorithm...");
    } else {
        // Perform actions for BFS algorithm
        console.log("Using BFS algorithm...");
    }
}