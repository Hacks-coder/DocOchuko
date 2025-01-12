<script lang="ts">
	let { children } = $props();
    const bgImageOne = "/img1.jpg"
</script>

<div style={`background-image: url("${bgImageOne}")`} class="bg-cover h-screen">
    <div class="h-screen ">
        {@render children()}
    </div>
</div>
