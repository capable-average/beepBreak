# Take Home Assignment

# **Drop Compliant Voicemails**

## The Problem

You’re a **software engineer at ClearPath Finance**, a company that helps people manage their credit card debt.

Every day, ClearPath’s system places **outbound phone calls** to customers to share important updates about their accounts.

One such customer is **Mike Rodriguez**.

Your system dials Mike’s number… but Mike is busy and doesn’t pick up.

Instead, the call goes to **voicemail**, and you hear this greeting:

> “Hi, you’ve reached Mike Rodriguez. I can’t take your call right now.
> 
> 
> Please leave your name, number, and a brief message after the beep.”
> 

### About Voicemail Greetings

These voicemail greetings are recorded by consumers, and in the real world they can be very different:

- Every person has a **different greeting**
- Some greetings **end with a beep**
- Some greetings **do not have a beep**
- Some greetings are **short**, others are **long**

---

## The Task

Your system now has an important job.

As soon as Mike’s greeting finishes, **ClearPath must leave a prerecorded, compliant voicemail message**, for example:

> “Hi, this is ClearPath Finance calling regarding your account.
> 
> 
> Please call us back at 800-555-0199. Thank you.”
> 

The challenge is deciding **exactly when to start playing this message**.

---

## What Does “Compliance” Mean?

**Anything the consumer hears must include:**

- The **company name**
- The **return phone number**

Anything the consumer **does not hear does not matter**. 

**⚠️ Important Note About the Beep:**

If a beep occurs, the consumer cannot hear anything spoken before the beep.

If the beep doesn’t occur, then the consumer can hear everything after their message ended.

This makes timing critical.

---

## Why Timing Matters

If your system:

- **Starts too early** → Mike might only hear
    
    *“…please call us back at 800-555-0199. Thank you”* ❌
    
    (Company name missing → non-compliant)
    
- **Starts too late** → The consumer may lose patience ❌

---

## Your Goal

> Develop a strategy to drop a compliant voicemail.
> 

---

## Input

- You are given **7 audio files** from calls that went to voicemail. You can access the audio files [here](https://drive.google.com/drive/folders/1RnRAkMxQTwsD5w3Kzd3BawpiLes2GOqn?usp=sharing).
- You must **stream** these audio files to simulate phone calls
    
    (Important: real phone calls are streaming, not pre-recorded chunks)
    

---

## Output

For **each audio file**, output the **timestamp(s)** at which you would start playing the voicemail.

---

## Hints (Optional)

- Speech-to-Text (e.g., Deepgram)
- LLMs for detecting common **end-of-greeting phrases**
- Handling cases with and without a beep

You are **encouraged to use AI**—both while writing code and while brainstorming your approach.

---

## What We’re Evaluating

This is a **learning-focused assignment**.

We are **not** expecting a perfect solution.

We care about:

- How you **think**
- How you **handle edge cases**
- How clearly you can **explain your logic**

---

## Deliverables

Please submit:

1. **Code**
    - Any programming language
2. **A short paragraph**
    - Explaining how your logic works
3. **A small demo**
    - Showing your solution in action

Please drop your submissions [here](https://forms.gle/XwAdpFPna7zkMXeC7).

---

## Final Note

This is the kind of real-world problem we solve every day for our customers.

There is no single correct answer.

If thinking through problems like this excites you, you’d be a great fit for our team, we’re excited to see your approach 🙂