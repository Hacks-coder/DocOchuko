<script lang="ts">
	import Footer from '../../components/static/Footer.svelte';
	import Header from '../../components/static/Header.svelte';
	import Hero from '../../components/static/Hero.svelte';

	let { children } = $props();
</script>
<Header />
<Hero />
{@render children()}
<Footer />
