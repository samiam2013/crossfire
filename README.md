# CrossFire

A dumb little project to get AI LLMs to fight one another

Name inspired by this youtube video of a show that ran on CNN until Jon Stewart destroyed the premise of the show, live. https://www.youtube.com/watch?v=aFQFB5YpDZE&t=2s

needs a `.env` in the root folder, there's an example. 

expects you to run it from the root folder of the repo using `go run .`

requires two keys, for openai and anthropic.

```env
OPENAI_API_KEY=yourkeyhere
ANTHROPIC_API_KEY=yourkeyhere
```

example output:

```text
➜  crossfire git:(main) ✗ go run .                  
welcome to crossfire


The topic of today's debate is: determinism is true rather than free will (whether or not it subjectively matters)


The first argument by claude for 'determinism is true rather than free will (whether or not it subjectively matters)' being true:  

Determinism is the idea that all events, including human actions and choices, are ultimately determined by prior events combined with the laws of nature. In a deterministic universe, if you could rewind time back to any point, events would always unfold the same way.

There are several compelling arguments for why determinism is likely true:

1. The laws of physics, as we currently understand them, are deterministic. The movement of all particles is governed by prior causes and fixed laws. Our brains are made of particles obeying these same deterministic laws.

2. The unconscious activity in our brains seems to precede and cause our conscious choices, rather than the other way around. This has been demonstrated in neuroscience experiments. What feels like a free choice is really just the output of prior unconscious processes.

3. Our genes, upbringing, circumstances, and experiences all shape our personality, preferences, and how we think and behave. We don't choose these formative factors. So how can our choices be free if they stem from aspects of ourselves we didn't choose?

4. In an indeterministic quantum universe, our decisions would be random rather than "free." True freedom of the will seems incompatible with both determinism and indeterminism.

5. All of our choices have prior causes, in a long unbroken chain. These causes ultimately trace back to events before we were born. How can we be ultimately responsible for our choices if the original causes lie outside us?

So in summary, the deterministic nature of physical laws, unconscious precursors to decisions, unchosen formative factors, problems with indeterminism, and the unbroken chain of prior causes all point to determinism being true and free will being an illusion. Let me know if you would like me to elaborate on any point!


The response argument by openai: Counterarguments to the case for determinism and in support of free will (or a non-deterministic worldview) include:

1. **Quantum Indeterminacy**: Quantum mechanics demonstrates that on the atomic and subatomic levels, events do not occur with a determinacy, but rather with probabilities. This indeterminacy at the micro level could allow for non-determined possibilities in the brain's functioning, opening a window for free will.

2. **Consciousness and Free Will**: The nature and origin of consciousness are not fully understood, and it's possible that consciousness plays a role in decision-making that is not strictly determined by physical laws. This could mean that our subjective experience of making choices is a genuine expression of free will.

3. **Complex Systems**: The brain is a highly complex system that might operate on principles of emergent properties, where the whole is more than the sum of its parts. In such systems, deterministic predictions based on component parts (such as individual neurons) might not accurately predict the system's behavior (such as thoughts and decisions), allowing for free will.

4. **Compatibilism**: Some philosophers argue that free will is compatible with determinism. According to compatibilism, free will means being able to act according to one's desires and motivations, without external constraints, even if those desires are determined by prior causes.

5. **Moral Responsibility and Societal Function**: The belief in free will supports the basis of moral responsibility. Without a belief in free will, concepts of justice and personal accountability may lose their grounding. This doesn't prove free will per se but argues for its practical necessity and perhaps hints at an aspect of the human condition that deterministic models don't fully capture.

6. **Incomplete Understanding of Physics**: Our current understanding of physics is incomplete. Theories of quantum gravity and the unification of general relativity with quantum mechanics are still works in progress. It's possible that a more complete theory of physics could provide a framework that supports the existence of free will.

7. **Unchosen Formative Factors vs. Active Processing**: While we don't choose our genes or early life experiences, the human capacity for reflection, learning, and adaptation allows us to transcend our initial conditions. We can critically assess our desires, change our goals, and act in ways that are not simply the direct results of past events.

In summary, the debate between determinism and free will involves deeply complex issues regarding physics, consciousness, morality, and the nature of human existence. The existence of free will may not be easily dismissible despite the persuasive arguments made by determinism.


The response argument by claude:  
The strongest arguments against determinism are:
1. Quantum indeterminacy indicates that physics is probabilistic, not deterministic, at its core. This opens up space for free will.
2. Consciousness is not fully understood and may allow for non-determined choices. Subjective experience of free will could be valid.  
3. As a complex emergent system, the brain's high-level behavior (thoughts, choices) may not be predictable from low-level deterministic processes.
4. Even in a determinist world, we can define free will compatibly as the ability to act on our desires without external constraint.
5. Belief in free will may be necessary for moral responsibility and social function. This suggests determinism misses key aspects of the human condition.
6. Incomplete physics, e.g. quantum gravity, could discover principles permitting free will. Current deterministic models may be wrong.
7. Humans can transcend their genes and background through reason & willpower. We aren't just results of factors outside our control.
So determinism has some strong arguments, but the case isn't settled. There remain significant reasons to doubt that all human choices are fully determined. The truth may be more complex than determinism or libertarian free will.


The response argument by openai: Rebutting the arguments against determinism, in favor of a deterministic or compatible view:

1. **Quantum Indeterminacy and Macroscopic Determinism**: While quantum mechanics introduces indeterminacy at the microscopic level, there is little evidence to suggest this randomness directly influences thought and decision-making in the structured, deterministic environment of the brain. Complex biological systems like the brain operate at a macroscopic level where classical physics' determinism effectively rules, making quantum indeterminacy unlikely to contribute significantly to free will.

2. **Consciousness and Determined Processes**: The lack of full understanding of consciousness does not inherently validate non-determined choices. Conscious experience could be the emergent result of complex but deterministic interactions between neurons. This means our subjective experience of free will could be an illusion created by these interactions, without contradicting deterministic principles.

3. **Predictability in Complex Systems**: The premise that a complex system like the brain could produce unpredictable, non-determined outcomes from determined processes potentially misunderstands chaos theory and emergent phenomena. Both can result from deterministic rules applied in complex systems, suggesting high-level behavior could still be determined by low-level deterministic processes.

4. **Compatibilism**: The compatibilist definition of free will (freedom to act according to one's determined desires) does not negate determinism but rather offers a reformulation of free will that is compatible with it. It suggests our actions, even if determined, can still be considered "free" in the sense that they align with our desires and motivations.

5. **Moral Responsibility and Determinism**: The belief in moral responsibility and social function does not necessitate the existence of libertarian free will. A deterministic framework can still accommodate concepts of responsibility and ethics by focusing on predictable outcomes of actions and the practical necessity of promoting societal well-being, rather than on an indeterminate notion of choice.

6. **Incomplete Physics as Irrelevant to Determinism in Human Behavior**: The potential for future physics discoveries (like quantum gravity) does not inherently imply that these will favor free will. Even if current deterministic models are incomplete, they still provide a robust framework for understanding natural phenomena, including human behavior, without requiring the introduction of free will as an explanatory variable.

7. **Transcendence Through Reason and Willpower as Determined**: The ability of humans to reflect, reason, and seemingly transcend their genetic and environmental backgrounds can itself be understood through deterministic processes. Learning, adaptation, and self-improvement can be explained by brain plasticity and environmental interactions, all of which can be influenced by previous determined states.

In summary, while arguments against determinism raise important points, they don't conclusively rule out a deterministic framework or necessarily validate the existence of free will in its libertarian sense. Determinism remains a robust way to understand human behavior and the universe, compatible with current scientific understanding and capable of incorporating the complexities of human thought and morality.



```
