<div>
    <h1>Here is the main page of the app</h1>
    <section>
        <button on:click="{handleClick}">Click This</button>
    </section>
</div>

<script>
    export let alex;

    const handleClick = (event) => {
        alert(`Your name is ${alex.Name} and age is ${alex.Age}.`);
    }
</script>

<style>
    h1, section {
        text-align: center;
    }

    button {
        border: none;
        padding: 10px;
        font-size: 1em;
        background-color: plum;
        color: white;
        border-radius: 4px;
    }
</style>